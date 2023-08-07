import pandas as pd
from sklearn.cluster import KMeans
from sklearn.metrics import silhouette_score
from scipy.spatial import ConvexHull
import geojson
import json



# Step 1: Read address data from the CSV file
address_df = pd.read_csv('area.csv')

# Step 2: Group addresses by area
# {
#   "type": "FeatureCollection",
#   "features": [
#     {
#       "type": "Feature",
#       "properties": {},
#       "geometry": {
#         "coordinates": [
#           90.38922321692223,
#           23.75292316348991
#         ],
#         "type": "Point"
#       }
#     },
#     {
#       "type": "Feature",
#       "properties": {},
#       "geometry": {
#         "coordinates": [
#           90.38979184523299,
#           23.753492722611696
#         ],
#         "type": "Point"
#       }
#     }
#   ]
# }

area_geojson = {
    "type": "FeatureCollection",
    "features": []
}


# Step 3: Create GeoJSON features for each area
for lat, lon , area in address_df:
    print(lat, lon , area)
   
