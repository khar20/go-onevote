# -*- coding: utf-8 -*-
"""
Created on Sat Oct 26 17:33:52 2024

@author: aries
"""

import face_recognition
import cv2
import json
import numpy as np
import os

#Definir los colores
color_amarillo = (0, 255, 255) #Color amarillo para desconocidos
color_rojo = (50, 50, 255) #Color rojo para conocidos pero sin parpadeo
color_verde = (125, 220, 0) #Color verde para conocidos con parpadeo (sí pasaron la prueba de antispoofing)
color_blanco = (255, 255, 255) #Color blanco para el color del texto

#Cargar características desde el archivo JSON
with open('notebooks/archivos_notebooks/caracteristicas.json', 'r') as f:
    data = json.load(f)

nombres = data["nombres"]

#Funcional
caracteristicas_conocidas = [encoding for encoding in data["caracteristicas"].values()]
# Se convierte las listas de características a numpy arrays
caracteristicas_conocidas = [np.array(encoding) for encoding in caracteristicas_conocidas]

# Iniciar la captura de video
video_capture = cv2.VideoCapture(1, cv2.CAP_DSHOW)

def calcular_ratio_ojos(landmarks):
    # Obtener los puntos de referencia de los ojos
    ojo_izquierdo = np.array(landmarks['left_eye'])
    ojo_derecho = np.array(landmarks['right_eye'])
    
    # Calcular la relación de aspecto
    alto_izquierdo = np.linalg.norm(ojo_izquierdo[1] - ojo_izquierdo[5]) + np.linalg.norm(ojo_izquierdo[2] - ojo_izquierdo[4])
    ancho_izquierdo = np.linalg.norm(ojo_izquierdo[0] - ojo_izquierdo[3])
    
    alto_derecho = np.linalg.norm(ojo_derecho[1] - ojo_derecho[5]) + np.linalg.norm(ojo_derecho[2] - ojo_derecho[4])
    ancho_derecho = np.linalg.norm(ojo_derecho[0] - ojo_derecho[3])

    ratio_izquierdo = alto_izquierdo / (2.0 * ancho_izquierdo)
    ratio_derecho = alto_derecho / (2.0 * ancho_derecho)

    return (ratio_izquierdo, ratio_derecho)

while True:
    # Captura un fotograma del video
    ret, frame = video_capture.read()
    if ret == False: break
    frame = cv2.flip(frame, 1) #Voltear la imagen para el efecto espejo

    # Encuentra todas las caras y sus características en el fotograma
    caras_en_frame = face_recognition.face_locations(frame) #Con el modelo por defecto
    caras_encodings_en_frame = face_recognition.face_encodings(frame, caras_en_frame)
    #Agregado: Para los ojos (antispoofing)
    landmarks_en_frame = face_recognition.face_landmarks(frame, caras_en_frame)    

    # Iterar sobre cada cara detectada en el fotograma
    for (top, right, bottom, left), face_encoding, landmarks in zip(caras_en_frame, caras_encodings_en_frame, landmarks_en_frame):
        # Verifica si la cara es conocida
        # Según la documentación, primero va la lista de rostros conocidos, y luego los rostros a comparar
        coincidencias = face_recognition.compare_faces(caracteristicas_conocidas, face_encoding)
        nombre = "Desconocido"
        color = color_amarillo #Color amarillo para desconocidos

        # Si hay coincidencias, usar el primer nombre encontrado
        if True in coincidencias:
            primer_coincidencia_index = coincidencias.index(True)
            nombre = nombres[primer_coincidencia_index]
            color = color_verde #Color verde

        #Calcular la relación de aspecto de los ojos
        ratio_izquierdo, ratio_derecho = calcular_ratio_ojos(landmarks)
        umbral_parpadeo = 0.25 #Ajustar este valor según se necesite

        #Detectar parpadeo
        parpadeo_detectado = (ratio_izquierdo < umbral_parpadeo) or (ratio_derecho < umbral_parpadeo)

        #Si se detecta parpadeo, se dibuja un cuadro alrededor de la cara
        if parpadeo_detectado:
            # Dibujar un cuadro alrededor de la cara (color verde)   
            cv2.rectangle(frame, (left, top -30), (right, top), color, 2) #Cuadro para el texto  
            cv2.rectangle(frame, (left, top), (right, bottom), color, 2) #Cuadro para el rostro      
        else:
            color = color_rojo #Conocidos que no pasaron la prueba de antispoofing
            # Dibujar un cuadro alrededor de la cara (color rojo)          
            cv2.rectangle(frame, (left, top -30), (right, top), color, 2) #Cuadro para el texto  
            cv2.rectangle(frame, (left, top), (right, bottom), color, 2) #Cuadro para el rostro   

        cv2.putText(frame, nombre, (left, top - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.75, color_blanco, 1) #(255, 255, 255) para color blanco

    # Mostrar el resultado
    cv2.imshow('Video', frame)

    # Romper el ciclo al presionar 'q' o Escape
    k = cv2.waitKey(1)
    if (k & 0xFF == ord('q')) or (k == 27 & 0XFF):
        break

# Cerrar la captura de video
video_capture.release()
cv2.destroyAllWindows()

'''
import cv2
import face_recognition

image = cv2.imread("Images/FotoPablo.jpg")
face_loc = face_recognition.face_locations(image)[0]
#print("face_loc: ", face_loc)

face_image_encodings = face_recognition.face_encodings(image, known_face_locations=[face_loc])[0]
#Para el vector de características

print("face_image_encodings: ", face_image_encodings)

cv2.rectangle(image, (face_loc[3], face_loc[0]), (face_loc[1], face_loc[2]), (0, 255, 0))
cv2.imshow("Image", image)
cv2.waitKey(0)
cv2.destroyAllWindows() 


#Video streaming
cap = cv2.VideoCapture(0, cv2.CAP_DSHOW)

while True:
    ret, frame = cap.read()
    if ret == False: break
    frame = cv2.flip(frame, 1)
    
    #Detección facial
    face_locations = face_recognition.face_locations(frame) #Por defecto
    #Para un buen rendimiento con cnn es necesaria la aceleración en la GPU (a través de la biblioteca CUDA de NVidia). 
    #También se deberá habilitar el soporte CUDA al compilar dlib.
    #face_locations = face_recognition.face_locations(frame, model="cnn") #Para usar con cnn
    if face_locations != []:
        for face_location in face_locations:
            face_frame_encodings = face_recognition.face_encodings(frame, known_face_locations=[face_location])[0]
            #result = face_recognition.compare_faces([face_frame_encodings], face_image_encodings)
            #Según la documentación, primero va la lista de rostros conocidos, y luego los rostros a comparar
            result = face_recognition.compare_faces([face_image_encodings], face_frame_encodings) 
            print("Result: ", result)

            if result[0] == True:
                text = "Pablo"
                color = (125, 220, 0) #Color verde
            else:
                text = "Desconocido"
                color = (50, 50, 255) #Color rojo

            cv2.rectangle(frame, (face_location[3], face_location[2]), (face_location[1], face_location[2] + 30), color, -1)       
            cv2.rectangle(frame, (face_location[3], face_location[0]), (face_location[1], face_location[2]), color, 2)
            cv2.putText(frame, text, (face_location[3], face_location[2] + 20), 2, 0.7, (255, 255, 255), 1) #(255, 255, 255) para color blanco
    #Fin detección facial
    
    cv2.imshow("Frame (presiones Escape para salir)", frame)
    k = cv2.waitKey(1)
    if k == 27 & 0XFF:
        break
    
cap.release()
cv2.destroyAllWindows()
'''