import React from "react";
import {connect} from 'react-redux'
import Symptoms from './Symptoms'
// import {redirect from './action/page'
import config from "./helper/config";

class Disease extends React.Component {
    render() {
        // const {name, sourceId, pinyin, namezh} = this.state;
        const {dises, id} = this.props;
        const dise = dises.find(x => x.id === id)
        if (!dise) {
            return <div>Need fetch server data</div>
        }

        return (
            <div>
                <div className="sectitle">Disease</div>
                <div>Name: {dise.namezh}, {dise.name}</div>
                <div>Source: {dise.source}</div>

                <Symptoms />
            </div>
        )
    }
}

const mapState = state => {
    return {dises: state.dises}
}

export default connect(mapState, {})(Disease)