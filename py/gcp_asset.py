from google.cloud import asset_v1p5beta1

project_id = 'jetstack-paul'
asset_types = ["containerregistry.googleapis.com/Image"]
page_size = 10

project_resource = "projects/{}".format(project_id)
content_type = asset_v1p5beta1.ContentType.RESOURCE
client = asset_v1p5beta1.AssetServiceClient()

# Call ListAssets v1p5beta1 to list assets.
response = client.list_assets(
    request={
        "parent": project_resource,
        "read_time": None,
        "asset_types": asset_types,
        "content_type": content_type,
        "page_size": page_size,
    }
)

for asset in response:
    image = {
        "name": asset.resource.data['name'].split('@')[0],
        "digest": asset.resource.data['name'].split('@')[1]
    }
    if 'tags' in asset.resource.data.keys():
        image['tags'] = asset.resource.data['tags']
    print(image)
