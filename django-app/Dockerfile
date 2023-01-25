FROM python:3.10

WORKDIR /app

COPY . .

RUN pip3 install -r requirements.txt

EXPOSE 8081

RUN chmod +x /app/skripta.sh

ENTRYPOINT [ "/app/skripta.sh" ]