db.auth('root', 'password');
db.createCollection('chats');
db.chats.insert({ id: 1, title: 'Default chat' });
