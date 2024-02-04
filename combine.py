import glob
import gzip
import json
from xml.etree import ElementTree

import fitparse

ACTIVITIES_GLOB = "./strava_data/activities/*"
# 2^32 / 360 - https://gis.stackexchange.com/questions/122186/convert-garmin-or-iphone-weird-gps-coordinates
CONVERSION = 11930465


def main():
    coords: list[list[list[float]]] = []
    count = 0
    for f in glob.glob(ACTIVITIES_GLOB):
        if f.endswith("fit.gz"):
            if fit_result := get_fit_coords(f):
                coords.append(fit_result)
        else:
            coords.append(get_gpx_coords(f))
        count += 1
        if count % 10 == 0:
            print(f"processed: {count} activities")

    with open("result.geojson", "wb") as f:
        f.write(
            json.dumps(
                {
                    "type": "MultiLineString",
                    "coordinates": coords,
                },
                separators=(",", ":"),
            ).encode()
        )


def get_gpx_coords(path: str):
    root = ElementTree.parse(path).getroot()
    pts = root.findall(".//{http://www.topografix.com/GPX/1/1}trkpt")
    return [[float(pt.attrib["lon"]), float(pt.attrib["lat"])] for pt in pts]


def get_fit_coords(path: str) -> list[list[float]] | None:
    pts: list[list[float]] = []
    with gzip.open(path, "r") as f:
        fitfile = fitparse.FitFile(f.read())

        # skip virtual rides
        if file_ids := [m for m in fitfile.get_messages("file_id")]:
            if file_ids[0].get_value("manufacturer") == "zwift":
                print(f"skipping zwift ride: {path}")
                return

        for record in fitfile.get_messages("record"):
            lat = record.get_value("position_lat")
            long = record.get_value("position_long")
            if lat and long:
                pts.append([long / CONVERSION, lat / CONVERSION])
    return pts


if __name__ == "__main__":
    main()
