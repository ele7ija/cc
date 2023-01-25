from django.db import models

# Create your models here.

class User(models.Model):
    jmbg = models.CharField(max_length=200)
    ime = models.CharField(max_length=200)
    prezime = models.CharField(max_length=200)
    je_student = models.BooleanField()