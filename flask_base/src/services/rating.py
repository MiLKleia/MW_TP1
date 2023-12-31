import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.ratings import RatingSchema
from src.schemas.ratings import BaseRatingSchema
from src.models.http_exceptions import *


rating_url = "http://ratings-beta.edu.forestier.re/songs/"  # URL de l'API users (golang)

def get_song_ratings(id):
    response = requests.request(method="GET", url=rating_url+id+"/ratings")
    return response.json(), response.status_code

def get_song_rating_by_user(id, user):
    response = requests.request(method="GET", url=rating_url+id+"/ratings/"+user)
    return response.json(), response.status_code

def add_rating(id, uid, new_rating):
    # on récupère le schéma song pour la requête vers l'API users

    song = requests.request(method="GET", url="http://localhost:8083/music/"+id)
    json_song = song.json()
    artist = json_song['artist']

    rating_schema = BaseRatingSchema().loads(json.dumps(new_rating), unknown=EXCLUDE)
    rating_schema["artist"] = artist
    rating_schema["user_id"] = uid
    rating_schema["song_id"] =  id

    # on crée le rating côté API 
    response = requests.request(method="POST", url=rating_url+id+"/ratings", json=rating_schema)

    return response.json(), response.status_code

def modify_rating(id, rating_id, new_rating):

    rating_schema = BaseRatingSchema().loads(json.dumps(new_rating), unknown=EXCLUDE)

    # on envoie le nouveau rating côté API 
    response = requests.request(method="PUT", url=rating_url+id+"/ratings/"+rating_id, json=rating_schema)

    return response.json(), response.status_code

def get_rating_by_id(id, rating_id):

    response = requests.request(method="GET", url=rating_url+id+"/ratings/"+rating_id)
    return response.json(), response.status_code


def delete_rating_by_id(id, rating_id):
    response = requests.request(method="DELETE", url=rating_url+id+"/ratings/"+rating_id)
    return response.status_code



