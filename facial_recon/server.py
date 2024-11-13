"""
Facial Recognition Login System
"""

import base64
import face_recognition
import cv2
import json
import numpy as np
import os
from flask import Flask, request, jsonify, session
from flask_cors import CORS

# Configurar la aplicación Flask
app = Flask(__name__)
CORS(app) #Habilitar CORS para toda la aplicación
#CORS(app, resources={r"/login": {"origins": "http://localhost:8080"}}) #Limitar CORS solo a la ruta /login
app.secret_key = 'super_secret_key'  # Se cambia esto por una clave segura en producción

'''
# Definir colores
color_amarillo = (0, 255, 255)
color_rojo = (50, 50, 255)
color_verde = (125, 220, 0)
color_blanco = (255, 255, 255)

# Cargar características de JSON
with open('notebooks/archivos_notebooks/caracteristicas.json', 'r') as f:
    data = json.load(f)
nombres = data["nombres"]
caracteristicas_conocidas = [np.array(encoding) for encoding in data["caracteristicas"].values()]

# Función para calcular el ratio de parpadeo
def calcular_ratio_ojos(landmarks):
    ojo_izquierdo = np.array(landmarks['left_eye'])
    ojo_derecho = np.array(landmarks['right_eye'])
    alto_izquierdo = np.linalg.norm(ojo_izquierdo[1] - ojo_izquierdo[5]) + np.linalg.norm(ojo_izquierdo[2] - ojo_izquierdo[4])
    ancho_izquierdo = np.linalg.norm(ojo_izquierdo[0] - ojo_izquierdo[3])
    alto_derecho = np.linalg.norm(ojo_derecho[1] - ojo_derecho[5]) + np.linalg.norm(ojo_derecho[2] - ojo_derecho[4])
    ancho_derecho = np.linalg.norm(ojo_derecho[0] - ojo_derecho[3])
    return alto_izquierdo / (2.0 * ancho_izquierdo), alto_derecho / (2.0 * ancho_derecho)

# Función para reconocimiento facial y autenticación
def verificar_usuario():
    video_capture = cv2.VideoCapture(1, cv2.CAP_DSHOW)
    usuario_autenticado = False

    while True:
        ret, frame = video_capture.read()
        if not ret: break
        frame = cv2.flip(frame, 1)
        caras_en_frame = face_recognition.face_locations(frame)
        caras_encodings_en_frame = face_recognition.face_encodings(frame, caras_en_frame)
        landmarks_en_frame = face_recognition.face_landmarks(frame, caras_en_frame)

        for (top, right, bottom, left), face_encoding, landmarks in zip(caras_en_frame, caras_encodings_en_frame, landmarks_en_frame):
            coincidencias = face_recognition.compare_faces(caracteristicas_conocidas, face_encoding)
            if True in coincidencias:
                nombre = nombres[coincidencias.index(True)]
                ratio_izquierdo, ratio_derecho = calcular_ratio_ojos(landmarks)
                umbral_parpadeo = 0.25
                if (ratio_izquierdo < umbral_parpadeo) or (ratio_derecho < umbral_parpadeo):
                    usuario_autenticado = True
                    session['username'] = nombre
                    video_capture.release()
                    cv2.destroyAllWindows()
                    return nombre
        cv2.waitKey(1)
    return None
'''
#Cargar las caracteristicas y los nombres
#with open('detector_facial/archivos_notebooks/caracteristicas.json', 'r') as f:
with open('archivos_notebooks/caracteristicas.json', 'r') as f:
    data = json.load(f)
nombres = data["nombres"]
caracteristicas_dict = data["caracteristicas"]

#Obtener caracteristicas del usuario
def obtener_caracteristicas_usuario(cip):
    for nombre, caracteristicas in caracteristicas_dict.items():
        if nombre == cip:
            return np.array(caracteristicas)
    return None

# Ruta para el login
@app.route('/login', methods=['POST'])
def login():
    cip = request.form.get('cip')
    imagen_b64 = request.form.get('captured-image')
    if not cip or not imagen_b64:
        return jsonify({"mensaje": "Falta el CIP o la imagen"}), 400
    
    print(f'CIP: {cip}. \nImagen_b64: {imagen_b64}')

    # Decodificar imagen
    imagen_data = np.frombuffer(base64.b64decode(imagen_b64), np.uint8)
    frame = cv2.imdecode(imagen_data, cv2.IMREAD_COLOR)

    # Obtener características del usuario
    caracteristicas_usuario = obtener_caracteristicas_usuario(cip)
    if caracteristicas_usuario is None:
        return jsonify({"mensaje": "Usuario no encontrado"}), 404

    # Detectar caras en la imagen capturada
    caras_en_frame = face_recognition.face_locations(frame)
    caras_encodings_en_frame = face_recognition.face_encodings(frame, caras_en_frame)

    for face_encoding in caras_encodings_en_frame:
        coincidencia = face_recognition.compare_faces([caracteristicas_usuario], face_encoding)[0]
        if coincidencia:
            return jsonify({"mensaje": "Usuario autenticado", "autenticado": True}), 200

    return jsonify({"mensaje": "Autenticación fallida", "autenticado": False}), 401


if __name__ == "__main__":
    app.run(port=5000, debug=True)
