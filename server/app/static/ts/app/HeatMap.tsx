import Map, {Source, Layer, LineLayer} from 'react-map-gl';
import React from "react";
import 'mapbox-gl/dist/mapbox-gl.css';

export const HeatMap = () => {
  const heatmapLayer: LineLayer = {
      'id': 'heatmap',
      'type': 'line',
      'source': 'mapbox-terrain',
      'source-layer': 'result',
      'layout': {
          'line-join': 'round',
          'line-cap': 'round'
      },
      'paint': {
          'line-color': '#EB9360',
          'line-width': 1,
          'line-opacity': 0.8
      },
  }
  return (
    <div>
      <Map
        mapLib={import('mapbox-gl')}
        initialViewState={{
          longitude: -73.95551,
          latitude: 40.73932,
          zoom: 10.74,
        }}
        style={{width: "100%", height: 1000}}
        mapStyle="mapbox://styles/mbutterfield/clt9fms1l003l01qqfjccc3i3"
        mapboxAccessToken={process.env.MAPBOX_ACCESS_TOKEN}
      >
          <Source id="heatmap" type="vector" url="mapbox://mbutterfield.8n1uc2fp">
            <Layer {...heatmapLayer} />
          </Source>
      </Map>
    </div>
  );
};
