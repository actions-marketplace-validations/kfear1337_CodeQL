from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'mysql://root:password@localhost/blogdb'
db = SQLAlchemy(app)

class Blog(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(80), nullable=False)
    content = db.Column(db.String(120), nullable=False)

    def __repr__(self):
        return '<Blog %r>' % self.title

@app.route('/search', methods=['GET'])
def search_blog():
    query = request.args.get('q')
    blogs = Blog.query.filter(Blog.title.like(f'%{query}%')).all()
    return jsonify([blog.title for blog in blogs])

if __name__ == '__main__':
    app.run(debug=True)
