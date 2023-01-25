from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
    path('<int:id>/', views.user, name='user'),
    path('register-student/', views.register_student, name='register_student'),
    path('register-prof/', views.register_prof, name='register_prof'),
]