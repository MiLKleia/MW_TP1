import json
import yaml
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError
from flask_login import login_user, logout_user, login_required, current_user


from src.models.http_exceptions import *
from src.schemas.music import SongUpdateSchema
from src.schemas.ratings import RatingAddUpdateSchema
from src.schemas.errors import *
import src.services.music as music_service
import src.services.rating as rating_service

# from routes import music
music = Blueprint(name="music", import_name=__name__)


@music.route('/<id>', methods=['GET'])
@login_required
def get_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """
    return music_service.get_song(id)


@music.route('/<id>', methods=['PUT'])
@login_required
def put_soong(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      requestBody:
        required: true
        content:
            application/json:
                schema: UserUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - users
    """
    # parser le body
    try:
        song_update = SongUpdateSchema().loads(json_data=json.dumps(yaml.load(request.data.decode('utf-8'))))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification de la musique (name, artist, album)
    try:
        return music_service.modify_song(id, song_update)
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")


@music.route('/', methods=['GET'])
@login_required
def get_songs():
    """
    ---
    get:
      description: Getting a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """
    return music_service.get_songs()


@music.route('/', methods=['POST'])
@login_required
def add_song():
    """
    ---
    get:
      description: Getting a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """

    try:
        song_to_add = SongUpdateSchema().loads(json_data=json.dumps(yaml.load(request.data.decode('utf-8'))))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification de la musique (name, artist, album)
    try:
        return music_service.add_song(song_to_add)
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
    



@music.route('/album/<title>', methods=['GET'])
@login_required
def get_album(title):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """

    return music_service.get_album(title)

@music.route('/artist/<name>', methods=['GET'])
@login_required
def get_artist(name):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """

    return music_service.get_artist(name)



@music.route('/<id>', methods=['DELETE'])
@login_required
def delete_song(id):
    
    return music_service.delete_song(id)

@music.route('/album/', methods=['GET'])
@login_required
def get_list_album():
    
    return music_service.get_list_album()

@music.route('/artist/', methods=['GET'])
@login_required
def get_list_artist():
    
    return music_service.get_list_artist()


@music.route('/<id>/ratings', methods=['GET'])
@login_required
def get_song_ratings(id):
    return rating_service.get_song_ratings(id)

@music.route('/<id>/ratings', methods=['POST'])
@login_required
def add_rating(id):
     
     try:
        rating_to_add = RatingAddUpdateSchema().loads(json_data=json.dumps(yaml.load(request.data.decode('utf-8'))))
     except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # ajout du rating (id, rating)
     try:
        return rating_service.add_rating(id, current_user.id, rating_to_add)
     except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
     except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
     

@music.route('/<id>/ratings/<rating_id>', methods=['PUT'])
@login_required
def modify_rating(id, rating_id):
     
     try:
        rating_to_change = RatingAddUpdateSchema().loads(json_data=json.dumps(yaml.load(request.data.decode('utf-8'))))
     except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification
     try:
        return rating_service.modify_rating(id, rating_id, rating_to_change)
     except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
     except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
     




@music.route('/<id>/ratings/<rating_id>', methods=['GET'])
@login_required
def get_rating_by_id(id, rating_id):
  
     try:
        return rating_service.get_rating_by_id(id, rating_id)
     except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
     except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
     
@music.route('/<id>/ratings/<rating_id>', methods=['DELETE'])
@login_required
def delete_rating_by_id(id, rating_id):
  
     try:
        return rating_service.delete_rating_by_id(id, rating_id)
     except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
     except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
     



