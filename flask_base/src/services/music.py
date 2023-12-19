import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.music import SongSchema
from src.models.http_exceptions import *


songs_url = "http://localhost:8083/music/"  # URL de l'API music (golang)

def get_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code

def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code


def add_song(song_register):
    
    music_model = MusicModel(song_register)
    # on récupère le schéma song pour la requête vers l'API users
    song_schema = SongSchema().loads(json.dumps(song_register), unknown=EXCLUDE)

    # on crée la musique côté API music
    response = requests.request(method="POST", url=songs_url, json=song_schema)

    return response.json(), response.status_code


def modify_song(id, song_update):
   
    # s'il y a quelque chose à changer côté API (name, artist, album)
    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)
    response = None
    if not SongSchema.is_empty(song_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=songs_url+id, json=song_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    return (response.json(), response.status_code) if response else get_song(id)

