import geojson
import h3
from shapely.geometry import shape, mapping,Polygon

# Input GeoJSON (same as before)
input_geojson = '''
{
    "type": "FeatureCollection",
    "name": "Dhaka GeoJSON v1",
    "crs": {
        "type": "name",
        "properties": {
            "name": "urn:ogc:def:crs:OGC:1.3:CRS84"
        }
    },
    "features": [
        {
            "type": "Feature",
            "properties": {
                "level_0": "Bangladesh",
                "level_1": "Dhaka",
                "level_2": "Dhaka",
                "level_3": "Jatrabari",
                "level_4": "Matuail",
                "level_5": "Green Model Town",
                "Area": "Mugda",
                "layer": "Dhaka Shapefile v1.2",
                "path": "/Users/pathaoltd/Downloads/GEO SPEKTRON/Shapefile Conversion/Shapefile Converted to GeoJSON/new one/Dhaka Shapefile v1.2/Dhaka Shapefile v1.2.shp"
            },
            "geometry": {
                "type": "MultiPolygon",
                "coordinates": [
                    [
                        [
                            [
                                90.446739427413704,
                                23.723267322656337
                            ],
                            [
                                90.446830749511804,
                                23.723279953002983
                            ],
                            [
                                90.446365356445384,
                                23.726999282836921
                            ],
                            [
                                90.449592590332003,
                                23.730651855468885
                            ],
                            [
                                90.456153869628963,
                                23.733119964599517
                            ],
                            [
                                90.458981000871518,
                                23.733695132909112
                            ],
                            [
                                90.459764822759837,
                                23.727659369793809
                            ],
                            [
                                90.455581659553829,
                                23.727528816753779
                            ],
                            [
                                90.453109790386577,
                                23.727485299044599
                            ],
                            [
                                90.453395006059694,
                                23.723829759616876
                            ],
                            [
                                90.446739427413704,
                                23.723267322656337
                            ]
                        ]
                    ]
                ]
            }
        }
    ]
}
'''

# Load GeoJSON
geojson_data = geojson.loads(input_geojson)

# Function to find central lat/lon of an H3 index
def h3_index_centroid(h3_index):
    lat, lon = h3.h3_to_geo(h3_index)
    return lat, lon

# Output GeoJSON
output_geojson = {
    "type": "FeatureCollection",
    "features": []
}

# Convert MultiPolygon to H3 indexes and find central lat/lon of each H3 index
for feature in geojson_data["features"]:
    if feature["geometry"]["type"] == "MultiPolygon":
        # Flatten the GeoJSON coordinates to a 2D list
        coordinates = feature["geometry"]["coordinates"]
        flattened_coordinates = []
        for polygon_coords in coordinates:
            for lon_lat in polygon_coords[0]:
                flattened_coordinates.append((lon_lat[0], lon_lat[1]))

        # Create a shapely Polygon
        polygon = Polygon(flattened_coordinates)

        # Convert the shapely Polygon to a GeoJSON dictionary
        polygon_geojson = mapping(polygon)

        # Get the H3 indexes that intersect with the polygon
        h3_indexes = list(h3.polyfill(polygon_geojson, res=13))

        for h3_index in h3_indexes:
            properties = feature["properties"]
            properties["h3_index"] = h3_index
            output_feature = geojson.Feature(geometry=h3.h3_to_geo_boundary(h3_index), properties=properties)
            output_geojson["features"].append(output_feature)

# Save the output GeoJSON to a file
with open('lat.geojson', 'w') as file:
    geojson.dump(output_geojson, file)

print("GeoJSON file 'output.geojson' created successfully.")
