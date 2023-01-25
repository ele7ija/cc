from django.shortcuts import render

# Create your views here.

from django.http import HttpResponse
from django.template import loader
from .models import User
import requests
import os

def sluzba():
    return os.getenv("FAKULTET")

def index(request):
    users = User.objects.all()
    template = loader.get_template('users/index.html')
    context = {
        'users': users,
        'fakultet': sluzba(),
    }
    return HttpResponse(template.render(context, request))

def user(request, id):
    try:
        user = User.objects.get(pk=id)
    except User.DoesNotExist:
        raise Http404("User does not exist")
    return render(request, 'users/user.html', {'user': user})

def exists(jmbg):
    url = "http://" + os.getenv("UNS_HOST") + "/user/" + str(jmbg)
    response = requests.get(url)
    print(response)
    if response.status_code == 400:
        return True
    return False

def register_student(request):
    ime = request.POST['ime']
    prezime = request.POST['prezime']
    jmbg = request.POST['jmbg']
    user = User(None, jmbg, ime, prezime, True)
    if exists(jmbg):
        template = loader.get_template('users/neuspesno.html')
        return HttpResponse(template.render({"user": user}, request))
    user.save()
    template = loader.get_template('users/uspesno.html')
    return HttpResponse(template.render({"user": user}, request))

def register_prof(request):
    ime = request.POST['ime']
    prezime = request.POST['prezime']
    jmbg = request.POST['jmbg']
    user = User(None, jmbg, ime, prezime, False)
    if exists(jmbg):
        template = loader.get_template('users/neuspesno.html')
        return HttpResponse(template.render({"user": user}, request))
    user.save()
    template = loader.get_template('users/uspesno.html')
    return HttpResponse(template.render({"user": user}, request))
