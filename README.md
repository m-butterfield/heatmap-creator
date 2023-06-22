# Heatmap Creator

Create your own heatmap by generating a geojson file from all your strava activities.
This file can then be loaded into any mapping software.
Currently supports `gpx` and `fit` files.

To run, update `ACTIVITIES_GLOB` in `combine.py` to the path of your data. Then run:

    poetry run python combine.py
