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

def add_rating(new_rating, id):
    # on récupère le schéma song pour la requête vers l'API users
    rating_schema = BaseRatingSchema().loads(json.dumps(new_rating), unknown=EXCLUDE)

    # on crée la musique côté API music
    response = requests.request(method="POST", url=rating_url+id+"/ratings", json=rating_schema)

    return response.json(), response.status_code

def add_rating(new_rating, id):
    # on récupère le schéma song pour la requête vers l'API users
    rating_schema = BaseRatingSchema().loads(json.dumps(new_rating), unknown=EXCLUDE)

    # on crée la musique côté API music
    response = requests.request(method="POST", url=rating_url+id+"/ratings", json=rating_schema)

    return response.json(), response.status_code

