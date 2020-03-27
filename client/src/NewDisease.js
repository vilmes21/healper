import React from "react";
import getAllSources from "./helper/getAllSources";
import {connect} from 'react-redux'
import { redirect } from './action/page'
import {addDise } from './action/dises'
import {getAll} from './action/source'
import config from "./helper/config";

class NewDisease extends React.Component {
    state = {
        sourceId: `3`,
        name: ``,
        pinyin: ``,
        namezh: ``
    }

    _change = ev => {
        const {name, value} = ev.target;
        this.setState({[name]: value})
    }

    getSourceNameById = id => {
        const x = this.props.source.find(x => x.id === id)

        if (!x) return `Err source`;
        
        return `${x.name} by ${x.author}`;
    }

    _submit = async ev => {
        ev.preventDefault();
      
        const data = {...this.state, sourceId: parseInt(this.state.sourceId)}

        try {
            const r = await fetch(config.host + config.apiUrl.createDisease, {
                method: 'POST',
                // headers: {
                //     'Content-Type': 'application/json'
                // },
                body: JSON.stringify(data)
            });
            const d = await r.json();
            console.log('fetch d:', d);

            if (!d.ok) return;

            const {addDise, redirect} =this.props;
            
            addDise({
                ...data,
                id: d.id,
                source: this.getSourceNameById(data.sourceId) 
            });

            redirect(config.urls.viewDise + d.id);
        } catch (e) {
            console.error('fetch Error:', e);
        }

        return false;
    }

    componentDidMount = async() => {
        const d = await getAllSources();
        
        if (!d) 
            return;
        
        //todo: add err notification

        const {getAll} = this.props;
        getAll(d);
    }

    render() {
        const {name, sourceId, pinyin, namezh} = this.state;
        const {source} = this.props;

        return (
            <div>
                <form onSubmit={this._submit}>
                    <input
                        className="ipt"
                        placeholder="Disease name"
                        type="text"
                        name="namezh"
                        value={namezh}
                        onChange={this._change}/>

                    <input
                        className="ipt"
                        placeholder="Pinyin"
                        type="text"
                        name="pinyin"
                        value={pinyin}
                        onChange={this._change}/>

                    <input
                        className="ipt"
                        placeholder="Name in English"
                        type="text"
                        name="name"
                        value={name}
                        onChange={this._change}/> {source.length > 0 && <select
                        required
                        className="ipt"
                        name="sourceId"
                        value={sourceId}
                        onChange={this._change}>
                        <option value={0}>-select source-</option>
                        {source.map(x => <option key={x.id} value={x.id}>
                            {x.name}
                            by {x.author}</option>)
}
                    </select>
}

                    <button className="bt">Save</button>
                </form>
            </div>
        )
    }
}

const mapState = state => {
    return {source: state.source}
}

export default connect(mapState, {getAll, addDise, redirect})(NewDisease)