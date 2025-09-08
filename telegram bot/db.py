import mysql.connector
from loadenv import DB_USER, DB_HOST, DB_PASSWORD, DB_NAME, DB_PORT

DB_CONFIG = {
    "host": DB_HOST,
    "user": DB_USER,
    "password": DB_PASSWORD,
    "database": DB_NAME,
    "port": DB_PORT
}

def insert_food(name, info, price):

    conn = mysql.connector.connect(**DB_CONFIG)
    cursor = conn.cursor()

    query = """
        INSERT INTO foods (name, info, price, created_at, updated_at)
        VALUES (%s, %s, %s, NOW(), NOW())
    """
    cursor.execute(query, (name, info, price))
    conn.commit()
    inserted_id = cursor.lastrowid

    cursor.close()
    conn.close()

    print(f"Food inserted with ID: {inserted_id}")
    return inserted_id
