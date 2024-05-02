def calcular_media_recortada(valores):
    valores_ordenados = sorted(valores)
    valores_recortados = valores_ordenados[50:-50]
    media_recortada = sum(valores_recortados) / len(valores_recortados)

    return media_recortada

def leer_valores_desde_archivo(nombre_archivo):
    with open(nombre_archivo, 'r') as archivo:
        valores = [float(linea.strip()) for linea in archivo.readlines()]
    return valores

nombre_archivo = 'elapsed_times.txt'
valores = leer_valores_desde_archivo(nombre_archivo)
media_recortada = calcular_media_recortada(valores)
print("La media recortada es:", media_recortada)
