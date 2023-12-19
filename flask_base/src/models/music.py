#from flask_login import UserMixin
from werkzeug.security import generate_password_hash
from src.helpers import db


# modèle de donnée pour la base de donnée utilisateur
# vous pouvez lier les utilisateurs de cette API à ceux de la vôtre (Golang) avec leur ID ou leur username

#class Song(UserMixin, db.Model):
class Song(db.Model):
    __tablename__ = 'music'

    id = db.Column(db.String(255), primary_key=True)
    name = db.Column(db.String(255), nullable=False)
    artist = db.Column(db.String(255), nullable=False)
    album = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, name, artist, album):
        self.id = uuid
        self.name = name
        self.artist = artist
        self.album = album

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.name or self.name == "") and \
               (not self.artist or self.artist == "")\
                (not self.album or self.album == "")
    
    

    