import pandas as pd
from sklearn.cluster import KMeans
from sklearn.metrics import silhouette_score
from scipy.spatial import ConvexHull
import geojson
import json
import sys

def cluster_lat_lon(points, num_clusters):
    if len(points) == 1:
        # Special case for areas with only one data point
        return [0], points
    else:
        kmeans = KMeans(n_clusters=num_clusters, n_init=1, init='random')
        kmeans.fit(points)
        return kmeans.labels_, kmeans.cluster_centers_

def create_geojson_point(coord):
    return geojson.Point(coord.tolist())

def create_geojson_polygon(coords):
    return geojson.Polygon([coords.tolist()])

# Step 1: Read address data from the CSV file
address_df = pd.read_csv(sys.argv[1])

# Step 2: Group addresses by area and sub_area
grouped_addresses = address_df.groupby(['area', 'sub_area'])

area_geojson = {
    "type": "FeatureCollection",
    "features": []
}

# Step 3: Create GeoJSON features for each area and sub_area group
for (area, sub_area), group in grouped_addresses:
    lat_lon_data = group[['lon', 'lat']].values
    num_data_points = len(lat_lon_data)

    if num_data_points == 1:
        # Handle the case where the area has only one data point
        point_coord = lat_lon_data[0]
        feature = geojson.Feature(geometry=create_geojson_point(point_coord), properties={'Area': area, 'SubArea': sub_area})
        area_geojson["features"].append(feature)
    else:
        # Step 4: Determine the optimal number of clusters using k-means with silhouette score
        max_clusters = min(num_data_points - 1, 10)  # Limit max_clusters to avoid errors
        optimal_num_clusters = 2  # Initialize with a minimum of 2 clusters

        if num_data_points > 2:
            silhouette_scores = []
            for num_clusters in range(2, max_clusters + 1):
                labels, _ = cluster_lat_lon(lat_lon_data, num_clusters)
                silhouette_scores.append((num_clusters, silhouette_score(lat_lon_data, labels)))

            # Find the optimal number of clusters with the highest silhouette score
            optimal_num_clusters = max(silhouette_scores, key=lambda x: x[1])[0]

        # Cluster the data with the optimal number of clusters
        labels, centers = cluster_lat_lon(lat_lon_data, optimal_num_clusters)

        polygons = []
        for cluster_id in range(optimal_num_clusters):
            cluster_points = lat_lon_data[labels == cluster_id]
            if len(cluster_points) >= 3:  # Check if there are enough points to create a convex hull
                hull = ConvexHull(cluster_points)
                polygon = create_geojson_polygon(cluster_points[hull.vertices])
                polygons.append(polygon)
            # Handle clusters with less than three points if necessary

        # Create the GeoJSON feature with properties
        feature = geojson.Feature(geometry=geojson.MultiPolygon(polygons), properties={'Area': area, 'SubArea': sub_area})
        area_geojson["features"].append(feature)

# Save the GeoJSON to a file
with open(sys.argv[2], 'w') as file:
    json.dump(area_geojson, file)

print("GeoJSON file 'output.geojson' created successfully.")
