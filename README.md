# Heatmap Creator

Create your own heatmap by generating a geojson file from all your strava activities.
This file can then be loaded into any mapping software.
Currently supports `gpx` and `fit` files.

[Example result](https://felt.com/map/Heatmap-Pir0y7hgS19By8t9COdiAdQA?loc=40.73191,-73.91632,12.43z)

To download your data from Strava, log in on the web and go to [https://www.strava.com/account](https://www.strava.com/account).
Then click on 'Get Started' under 'Download or Delete Your Account', and you should see the 'Download Request (optional)' option.
Click 'Request Your Archive' and you should receive an email shortly with a link to download your data.

Once you have your data, unzip it and update `ACTIVITIES_GLOB` in `combine.py` to the path to the 'activities' folder.
Then run:

    poetry run python combine.py

Now you can use the output file as you wish: `result.geojson`
