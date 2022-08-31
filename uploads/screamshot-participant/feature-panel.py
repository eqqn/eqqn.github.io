from flask import Flask, request, jsonify, make_response, render_template, render_template_string,redirect, url_for, Response
from flask_sqlalchemy import SQLAlchemy
import uuid
from  werkzeug.security import generate_password_hash, check_password_hash
import jwt
import imgkit
from datetime import datetime, timedelta
from functools import wraps
import base64
import os
import logging
from logging.handlers import RotatingFileHandler
from time import strftime

app = Flask(__name__, template_folder='./template')

@app.route("/admin", methods=['POST'])
def checkadminapi():
    data = request.form
    uri, output = data.get('url'), data.get('output')
    print(uri)
    print(output)
    #print(data.get('wkhtmltoimagepath'))
    status_to_ret = ""
    err_ret = ""
    bool_hid=""
    # RCE
    if len(data.get('wkhtmltoimage')) > 2:
        print("1")
        param = data.get('wkhtmltoimage')
        print(param)
        try:
            config = imgkit.config(wkhtmltoimage=param)
            imgkit.from_url(uri, str('./static/'+output), config=config)
        except Exception as e:
            err_ret = str(e)
            bool_hid = "true"
            status_to_ret = ""
        # RCE END
    else:
        print("2")
        try:
            imgkit.from_url(uri, str('./static/'+output))
                # with open(str('./'+output), "rb") as f:
                #     dataimg = base64.b64encode(f.read())
                # parse = dataimg.decode('utf-8')
            parse = output
            print(parse)
                # os.remove(output)
            status_to_ret = str(parse)
            bool_hid = "false"
            err_ret = ""

        except Exception as e:
            err_ret = str(e)
            bool_hid = "true"
            status_to_ret = ""
        print(status_to_ret)
        print(bool_hid)
        return render_template("panel.html", RESULT=status_to_ret, ERROR_RESULT=err_ret)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=False)
