const API_URL = 'http://localhost:8080/todos';

document.addEventListener('DOMContentLoaded', () => {
    fetchTodos();

    document.getElementById('todo-form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const title = document.getElementById('title').value;

        const res = await fetch(API_URL, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title, done: false })
        });

        if (res.ok) {
            document.getElementById('title').value = '';
            fetchTodos();
        }
    });
});

async function fetchTodos() {
    const res = await fetch(API_URL);
    const todos = await res.json();

    const list = document.getElementById('todo-list');
    list.innerHTML = '';

    todos.forEach(todo => {
        const item = document.createElement('li');
        item.className = 'list-group-item d-flex justify-content-between align-items-center';

        item.innerHTML = `
            <span>${todo.done ? '<del>' + todo.title + '</del>' : todo.title}</span>
            <div>
                <button class="btn btn-sm btn-success me-2" onclick="toggleDone(${todo.id}, ${todo.done})">
                    ${todo.done ? 'Undo' : 'Selesai'}
                </button>
                <button class="btn btn-sm btn-danger" onclick="deleteTodo(${todo.id})">Hapus</button>
            </div>
        `;

        list.appendChild(item);
    });
}

async function toggleDone(id, currentStatus) {
    await fetch(`${API_URL}/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ done: !currentStatus })
    });
    fetchTodos();
}

async function deleteTodo(id) {
    await fetch(`${API_URL}/${id}`, {
        method: 'DELETE'
    });
    fetchTodos();
}
