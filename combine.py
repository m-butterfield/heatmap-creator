import glob
import json
from xml.etree import ElementTree

ACTIVITIES_GLOB = "/Users/matt/Google Drive/My Drive/strava data/activities/*.gpx"

ElementTree.register_namespace("", "http://www.topografix.com/GPX/1/1")


def main():
    coords: list[list[list[float]]] = []
    count = 0
    for f in glob.glob(ACTIVITIES_GLOB):
        root = ElementTree.parse(f).getroot()
        pts = root.findall(".//{http://www.topografix.com/GPX/1/1}trkpt")
        coords.append(
            [[float(pt.attrib["lon"]), float(pt.attrib["lat"])] for pt in pts]
        )
        count += 1
        if count % 100:
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


if __name__ == "__main__":
    main()
