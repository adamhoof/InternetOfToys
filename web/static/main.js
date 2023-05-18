// main.js


window.onload = function() {
    let toys = [];
    fetch('/api/rooms/' + roomId)
        .then(response => response.json())
        .then(room => {
            room.toys.forEach(toy => {
                fetch('/api/toys/' + toy.id + '/commands')
                    .then(response => response.json())
                    .then(commands => {
                        let toyItem = room.toys.find(t => t.id === toy.id);
                        if (toyItem) {
                            toyItem.commands = commands;
                        }
                    });
            });
            toys = room.toys;
        });
};

function showToyDetails(toys, toyId) {
    let toy = toys.find(t => t.id === toyId);
    document.getElementById('toy-name').innerText = toy.name;
    toy.commands.forEach(command => {
        let commandElement;
        if (command.type === 'button') {
            commandElement = document.createElement('button');
            commandElement.innerText = command.name;
            commandElement.onclick = function() {
                fetch('/api/toys/' + toyId + '/control', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ command: command.name })
                });
            };
        } else if (command.type === 'slider') {
            commandElement = document.createElement('input');
            commandElement.type = 'range';
            commandElement.oninput = function() {
                fetch('/api/toys/' + toyId + '/control', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ command: command.name, value: commandElement.value })
                });
            };
        }
        document.getElementById('toy-commands').appendChild(commandElement);
    });
}
