from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma music de sortie (renvoyé au front)
class SongSchema(Schema):
    id = fields.String(description="UUID")
    name = fields.DateTime(description="Name")
    artist = fields.String(description="Artist")
    album = fields.String(description="Album")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("artist") or obj.get("artist") == "") and \
               (not obj.get("album") or obj.get("album") == "")



# Schéma musique de modification (name, artist, album)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "") or
                ("artist" in data and data["artist"] != "") or
                ("album" in data and data["album"] != "")):
            raise ValidationError("at least one of ['name','artist','album'] must be specified")
