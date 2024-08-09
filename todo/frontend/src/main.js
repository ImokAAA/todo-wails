document.getElementById('add').addEventListener('click', function() {
    const newTodoText = document.getElementById('inp').value;
    const newTodoDateTime = document.getElementById('datetime-inp').value;

    if (newTodoText.trim() === '' || newTodoDateTime.trim() === '') return;

    const todoList = document.getElementById('todo-list');

    const li = document.createElement('li');
    li.className = 'undone';

    const taskText = document.createElement('span');
    taskText.textContent = newTodoText;

    const dateTime = document.createElement('span');
    dateTime.className = 'datetime';
    dateTime.textContent = `(${new Date(newTodoDateTime).toLocaleString()})`;

    taskText.appendChild(dateTime);

    // Create actions container
    const actions = document.createElement('div');
    actions.className = 'actions';

    // Create done/undone button
    const toggleDoneButton = document.createElement('button');
    toggleDoneButton.textContent = 'Done';
    toggleDoneButton.className = 'toggle-done';

    toggleDoneButton.addEventListener('click', function() {
        li.classList.toggle('done');
        li.classList.toggle('undone');
        toggleDoneButton.textContent = li.classList.contains('done') ? 'Undone' : 'Done';
    });

    // Create delete button
    const deleteButton = document.createElement('button');
    deleteButton.textContent = 'Delete';

    deleteButton.addEventListener('click', function() {
        if (confirm('Are you sure you want to delete this task?')) {
            todoList.removeChild(li);
        }
    });

    actions.appendChild(toggleDoneButton);
    actions.appendChild(deleteButton);

    li.appendChild(taskText);
    li.appendChild(actions);
    todoList.appendChild(li);

    document.getElementById('inp').value = '';
    document.getElementById('datetime-inp').value = '';
});
