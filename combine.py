import glob
from xml.etree import ElementTree

ACTIVITIES_GLOB = "/Users/matt/Google Drive/My Drive/strava data/activities/*.gpx"

START = """<?xml version="1.0" encoding="UTF-8"?>
<gpx creator="StravaGPX" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd" version="1.1" xmlns="http://www.topografix.com/GPX/1/1">
 <metadata>
  <time>1986-08-06-02T17:11:00Z</time>
 </metadata>"""
END = """
</gpx>"""

TRK_START = """
 <trk>
  <name>Heatmap</name>
  <type>9</type>
  <trkseg>
   """
TRK_END = """</trkseg>
 </trk>"""

ElementTree.register_namespace("", "http://www.topografix.com/GPX/1/1")


def main():
    trks = ""
    for f in glob.glob(ACTIVITIES_GLOB):
        root = ElementTree.parse(f).getroot()
        pts = root.findall(".//{http://www.topografix.com/GPX/1/1}trkpt")
        trkpts = "".join(
            [
                ElementTree.tostring(p)
                .decode("utf-8")
                .replace(' xmlns="http://www.topografix.com/GPX/1/1"', "")
                for p in pts
            ]
        )
        trks += TRK_START + trkpts + TRK_END
    with open("result.gpx", "wb") as f:
        f.write((START + trks + END).encode())


if __name__ == "__main__":
    main()
