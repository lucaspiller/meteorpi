#!/bin/bash

OUTDIR="/tmp/media-new"
mkdir $OUTDIR

nice -n 19 tail -42000 data/data.csv > /tmp/week.csv
nice -n 19 tail -3500 /tmp/week.csv > /tmp/day.csv

date -1d > /dev/null 2>&1
if [ "$?" == "0" ]; then
  YESTERDAY="date -v-1d"
else
  YESTERDAY="date -d 'yesterday'"
fi

nice -n 19 gnuplot <<EOF
set datafile separator ","
set term svg size 700,350

set timefmt "%Y-%m-%d %H:%M:%S +0000 UTC"
set xdata time
set format x "%H:%M"

set grid ytics lc rgb "#bbbbbb" lw 1 lt 0
set grid xtics lc rgb "#bbbbbb" lw 1 lt 0

set xrange [ system("$YESTERDAY '+%s'") : system("date '+%s'") ]

set output '$OUTDIR/pressure.svg'
set ylabel "Pressure (hPa)"
set yrange  [ 950 : 1050 ]
plot '<grep BMP180 /tmp/day.csv' using 1:4 with lines title 'Pressure' smooth csplines
unset yrange

set output '$OUTDIR/remote1-vcc.svg'
set ylabel "Battery (V)"
set yrange  [ 2.5 : 4.5 ]
plot '<grep Remote1 /tmp/day.csv' using 1:5 with lines title 'Battery' smooth csplines
unset yrange

set ytics nomirror
set y2tics
set ylabel "Temperature (C)"
set y2label "Humidity (RH)"
set yrange  [ -10 : 40 ]
set y2range [ 0 : 100 ]

set output '$OUTDIR/internal.svg'
plot '<grep DHT22 /tmp/day.csv' using 1:3 with lines title 'Temperature' smooth csplines, '<grep DHT22 /tmp/day.csv' using 1:4 with lines title 'Humidity' smooth csplines axes x1y2

set output '$OUTDIR/remote1-temp.svg'
d(t, h)=(t - ((100 - h) / 5))
plot '<grep Remote1 /tmp/day.csv' using 1:3 with lines title 'Temperature' smooth csplines, '<grep Remote1 /tmp/day.csv' using 1:4 with lines title 'Humidity' smooth csplines axes x1y2, '<grep Remote1 /tmp/day.csv' using 1:(d(column(3), column(4))) with lines title 'Dewpoint' smooth csplines axes x1y1
EOF

NOW=`date`

cat > $OUTDIR/index.html <<EOF
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="refresh" content="30">
    <style>
      body {
        font-family: Arial, sans;
        font-size: 14px;
      }
      h1 { 
        margin: 0;
        font-size: 14px;
      }
    </style>
  </head>
  <body>
    <div>
      Day | <a href="/week.html">Week</a>
    </div>
    <div>
      <h1>Inside</h1>
      <img src="internal.svg">
      <img src="pressure.svg">
    </div>

    <div>
      <h1>Outside</h1>
      <img src="remote1-temp.svg">
      <img src="remote1-vcc.svg">
    </div>

    <div>
      Generated $NOW
    </div>
  </body>
</html>
EOF

date -v-1d > /dev/null 2>&1
if [ "$?" == "0" ]; then
  YESTERDAY="date -v-7d"
else
  YESTERDAY="date -d '1 week ago'"
fi

nice -n 19 gnuplot <<EOF
set datafile separator ","
set term svg size 700,350

set timefmt "%Y-%m-%d %H:%M:%S +0000 UTC"
set xdata time
set format x "%d/%m"

set grid ytics lc rgb "#bbbbbb" lw 1 lt 0
set grid xtics lc rgb "#bbbbbb" lw 1 lt 0

set xrange [ system("$YESTERDAY '+%s'") : system("date '+%s'") ]

set output '$OUTDIR/week-pressure.svg'
set ylabel "Pressure (hPa)"
set yrange  [ 950 : 1050 ]
plot '<grep BMP180 /tmp/week.csv' using 1:4 with lines title 'Pressure' smooth csplines
unset yrange

set output '$OUTDIR/week-remote1-vcc.svg'
set ylabel "Battery (V)"
set yrange  [ 2.5 : 4.5 ]
plot '<grep Remote1 /tmp/week.csv' using 1:5 with lines title 'Battery' smooth csplines
unset yrange

set ytics nomirror
set y2tics
set ylabel "Temperature (C)"
set y2label "Humidity (RH)"
set yrange  [ -10 : 40 ]
set y2range [ 0 : 100 ]

set output '$OUTDIR/week-internal.svg'
plot '<grep DHT22 /tmp/week.csv' using 1:3 with lines title 'Temperature' smooth csplines, '<grep DHT22 /tmp/week.csv' using 1:4 with lines title 'Humidity' smooth csplines axes x1y2

set output '$OUTDIR/week-remote1-temp.svg'
d(t, h)=(t - ((100 - h) / 5))
plot '<grep Remote1 /tmp/week.csv' using 1:3 with lines title 'Temperature' smooth csplines, '<grep Remote1 /tmp/week.csv' using 1:4 with lines title 'Humidity' smooth csplines axes x1y2, '<grep Remote1 /tmp/week.csv' using 1:(d(column(3), column(4))) with lines title 'Dewpoint' smooth csplines axes x1y1
EOF

NOW=`date`

cat > $OUTDIR/week.html <<EOF
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="refresh" content="30">
    <style>
      body {
        font-family: Arial, sans;
        font-size: 14px;
      }
      h1 { 
        margin: 0;
        font-size: 14px;
      }
    </style>
  </head>
  <body>
    <div>
      <a href="/">Day</a> | Week
    </div>
    <div>
      <h1>Inside</h1>
      <img src="week-internal.svg">
      <img src="week-pressure.svg">
    </div>

    <div>
      <h1>Outside</h1>
      <img src="week-remote1-temp.svg">
      <img src="week-remote1-vcc.svg">
    </div>

    <div>
      Generated $NOW
    </div>
  </body>
</html>
EOF

mkdir /tmp/media
cp /tmp/media-new/* /tmp/media
rm -Rf /tmp/media-new
