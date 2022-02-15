from waitress import serve
from flask import Flask
import os
import time
import mysql.connector as database
import random



app = Flask(__name__)

@app.route('/')
def index():
    #get the current time in nanoseconds

    before = time.time_ns()

    try:
        conn = database.connect(
            user="root",
            password="root",
            host="db",
            database="benchmarks"
        )
    except database.Error as e:
        return "Error connecting to MariaDB Platform: " + str(e)

    # Get Cursor
    try:
        cur = conn.cursor()
    except database.Error as e:
        return "Error getting cursor: " + str(e)
    

    #generate a random number between 1 and 10
    random_number = random.randint(1, 10)
    if random_number == 1:
        useless = 0
        for i in range(1, 100000000):
            useless += i

    end = time.time_ns() - before

    try:
        cur.execute("INSERT INTO benchmarks (backType, execTime) VALUES (%s, %s)", ("python",end))
        conn.commit()
    except database.Error as e:
        return "Error : " + str(e)

    return "python " + str(end) + " ns"


if __name__ == '__main__':
    serve(app, port=8080, host='0.0.0.0')