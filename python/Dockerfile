FROM python:3.6.5-stretch

ENV APP_DIRECTORY /var/www/helloapp
ADD . $APP_DIRECTORY

RUN apt-get update && apt-get upgrade -y
RUN easy_install pip
RUN pip install pip==10.0.1
RUN pip install gunicorn==19.8.1

WORKDIR $APP_DIRECTORY

CMD gunicorn -b :8000 main:app
