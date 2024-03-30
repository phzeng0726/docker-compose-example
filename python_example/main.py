from flask import Flask, jsonify

app = Flask(__name__)

# 假設這是你的資料庫
tasks = [
    {
        'id': 1,
        'title': '買牛奶',
        'description': '去超市買一瓶牛奶',
        'done': False
    },
    {
        'id': 2,
        'title': '學習 Python',
        'description': '完成 Flask RESTful API 教學',
        'done': False
    }
]

# 獲取所有任務
@app.route('/api1/tasks', methods=['GET'])
def get_tasks():
    return jsonify({'tasks': tasks})


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
