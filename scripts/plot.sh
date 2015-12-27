#!/bin/bash

tail -10000 data/data.csv > data/recent.csv

date -v-1d 2>&1 > /dev/null
if [ "$?" == "0" ]; then
  YESTERDAY="date -v-1d"
else
  YESTERDAY="date -d 'yesterday'"
fi

gnuplot <<EOF
set datafile separator ","
set term svg

set timefmt "%Y-%m-%d %H:%M:%S +0000 UTC"
set xdata time
set format x "%H:%M"

set xrange [ system("$YESTERDAY '+%s'") : system("date '+%s'") ]

set output 'media/pressure.svg'
set ylabel "Pressure (mb)"
plot '<grep BMP180 data/recent.csv' using 1:4 with lines title 'Pressure' smooth csplines

set output 'media/remote1-vcc.svg'
set ylabel "Battery (V)"
set yrange  [ 2.5 : 4.5 ]
plot '<grep Remote1 data/recent.csv' using 1:5 with lines title 'Battery' smooth csplines
unset yrange

set ytics nomirror
set y2tics
set ylabel "Temperature (C)"
set y2label "Humidity (RH)"
set yrange  [ -10 : 40 ]
set y2range [ 0 : 100 ]

set output 'media/internal.svg'
plot '<grep DHT22 data/recent.csv' using 1:3 with lines title 'Temperature' smooth csplines, '<grep DHT22 data/recent.csv' using 1:4 with lines title 'Humidity' smooth csplines axes x1y2

set output 'media/remote1-temp.svg'
plot '<grep Remote1 data/recent.csv' using 1:3 with lines title 'Temperature' smooth csplines, '<grep Remote1 data/recent.csv' using 1:4 with lines title 'Humidity' smooth csplines axes x1y2
EOF

NOW=`date`

cat > media/index.html <<EOF
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="refresh" content="30">
    <style>
      body {
        font-family: Arial, sans;
      }
    </style>
  </head>
  <body>
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
