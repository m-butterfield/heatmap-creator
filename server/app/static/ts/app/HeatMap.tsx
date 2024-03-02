import Map from 'react-map-gl';
import React from "react";

export const HeatMap = () => {
  return (
    <div>
      <Map
        mapLib={import('mapbox-gl')}
        initialViewState={{
          longitude: -100,
          latitude: 40,
          zoom: 3.5
        }}
        style={{width: "100%", height: 1000}}
        mapStyle="mapbox://styles/mbutterfield/clt9fms1l003l01qqfjccc3i3"
        mapboxAccessToken={process.env.MAPBOX_ACCESS_TOKEN}
      />;
    </div>
  );
};
