import React from "react";
import Redirec from "./Redirec";

class NewDisease extends React.Component {
    state = {
        sourceId: 0,
        name: ``,
        nameZh: ``,
        sources: [
            {
                name: `Learn TCM`,
                author: "John Doe",
                id: 1
            }, {
                name: `TCM`,
                author: "Mark Tum",
                id: 2
            }
        ]
    }

    _change = ev => {
        const {name, value} = ev.target;
        this.setState({[name]: value})
    }

    _submit = ev => {
        ev.preventDefault();
        console.log(`submitting:`)
        console.table(this.state)
    }

    render() {
        const {name, sourceId, sources, nameZh} = this.state;

        return (
            <div>
                <form onSubmit={this._submit}>
                    <input
                        placeholder="Disease name"
                        type="text"
                        name="nameZh"
                        value={nameZh}
                        onChange={this._change}/>

                    <input
                        placeholder="Name in English"
                        type="text"
                        name="name"
                        value={name}
                        onChange={this._change}/>

                    <select name="sourceId" value={sourceId} onChange={this._change}>
                        <option value={0}>-select-</option>
                        {sources.map(x => <option key={x.id} value={x.id}>{x.author}
                            - {x.name}</option>)
}
                    </select>

                    <button>Save</button>
                </form>
            </div>
        )
    }
}

export default NewDisease;