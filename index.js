class Block {
    id = crypto.randomUUID();
    static name = "Unnamed";

    update(input) {
        alert("Input: " + input);
    }

    render() {
        const div = document.createElement("div");
        div.textContent = "empty div";
    }
}

class Heading extends Block {
    dom = document.createElement("h1");
    static name = "Heading";

    update(input) {
        this.dom.textContent += input;
    }

    render() {
        return this.dom;
    }
}

let board = null;
let masterContextMenu = null;

const blocks = [];
const tools = {};

document.addEventListener("DOMContentLoaded", (e) => {
    console.log("Started...");

    board = document.querySelector("main");
    masterContextMenu = document.querySelector("#master-context");

    modifyPrototypes();

    __driverCode__();

    createContextMenu(Object.values(tools));
});

function updateBoard() {
    blocks.forEach(b => {
        b.update("Heading 1")
        board.appendChild(b.render());
    });
}

function modifyPrototypes() {
    blocks.push = function () {
        const ret = Array.prototype.push.apply(this, arguments);
        updateBoard();
        return ret;
    }
}

function __driverCode__() {
    tools[Heading.name] = Heading;
}

function createContextMenu(tools) {
    tools.forEach((tool, idx) => {
        const item = document.createElement("div");
        item.classList.add("context-element");
        item.textContent = tool.name;
        item.addEventListener("click", (e) => {
            newTool(tool.name);
        });
        masterContextMenu.appendChild(item);

        if (idx != tools.length - 1) {
            const sep = document.createElement("div");
            sep.classList.add("context-seperator");
            masterContextMenu.appendChild(sep);
        }
    });
}

function newTool(name) {
    blocks.push(new tools[name]());
}