from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma music de sortie (renvoyé au front)
# on accepte un comment vide

class RatingSchema(Schema):
    id = fields.String(description="UUID")
    song_id = fields.String(description="UUID")
    user_id = fields.String(description="UUID")
    comment = fields.String(description="Comment")
    rating_date = fields.DateTime(description="rating date")
    rating = fields.Int(description="rating")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("song_id") or obj.get("song_id") == "") and \
               (not obj.get("user_id") or obj.get("user_id") == "") and \
               (not obj.get("rating_date") or obj.get("rating_date") == "") and \
               (not obj.get("rating") or obj.get("rating") == "")

class BaseRatingSchema(Schema):
    id = fields.String(description="UUID")
    comment = fields.String(description="Comment")
    rating = fields.Int(description="rating")
    artist = fields.String(description="Comment")
    song_id = fields.String(description="UUID")
    user_id = fields.String(description="UUID")


# Schéma musique de modification (name, artist, album)
class SongUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("rating" in data and data["rating"] != "")):
            raise ValidationError("all fields must be provided")
