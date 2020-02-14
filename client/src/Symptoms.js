import React from "react";
import Options1 from "./Options1";
import Cart from "./Cart";
import CandidateTxt from "./CandidateTxt";
import {connect} from 'react-redux'
import {redirect} from './action/page'
import config from "./helper/config";
import {fillOrgans} from "./action/organs";
import {fillSympAdjs} from "./action/sympAdjs";
import fetchCategory from "./helper/fetchCategory";

class Symptoms extends React.Component {
    state = {
        sympKeyword: ``,
        candidate: [],
        cart: [],
        options: []
    }

    componentDidMount = async() => {
        const {fillOrgans, fillSympAdjs} = this.props;

        const orgs = await fetchCategory(config.categoryType.organ);
        if (orgs) 
            fillOrgans(orgs);
        
        const symps = await fetchCategory(config.categoryType.sympAdj);
        if (symps) 
            fillSympAdjs(symps);
        }
    
    _change = ev => {
        const {name, value} = ev.target;
        this.setState({[name]: value});

        const {organs, sympAdjs} = this.props;
        const {candidate} = this.state;

        //if organ is already set, then allow only adjective next.
        const collection = candidate.length === 1
            ? sympAdjs
            : organs.concat(sympAdjs);
        const matched = collection.filter(x => x.pinyin.indexOf(value) > -1 || x.name.indexOf(value) > -1 || x.namezh.indexOf(value) > -1);
        // console.table(matched)
        this.setState({options: matched})
    }

    appendCandidate = x => {
        const {candidate} = this.state;
        if (candidate.length === 2) 
            return;
        
        const updater = (prevState, p) => {
            //means organ has been set. Now only allow adjective
            const newState = {
                candidate: [
                    ...prevState.candidate,
                    x
                ]
            };

            if (candidate.length === 0) {
                newState.options = prevState
                    .options
                    .filter(y => y.catTypeId === config.categoryType.sympAdj);
            }

            return newState;
        }

        this.setState(updater);
    }

    cartSymp = () => {
        const {candidate} = this.state;
        if (candidate.length === 0) 
            return;
        
        const updater = (prevState, p) => {
            const newPair = prevState.candidate.length === 1
                ? [config.catnum.noOrganId, prevState.candidate[0]]
                : [...prevState.candidate];

            return {
                cart: [
                    ...prevState.cart,
                    newPair
                ],
                candidate: []
            }
        }

        this.setState(updater);
    }

    render() {
        const {sympKeyword,cart, candidate, options} = this.state;
        return (
            <div>
                <div className="sectitle">Symptoms</div>

                <Cart objArr={cart} />

                <div>
                    <CandidateTxt objArr={candidate}/> {candidate.length > 0 && <button className="bt" onClick={this.cartSymp}>Add to cart</button>
}
                </div>

                <input
                    className="ipt"
                    onChange={this._change}
                    type="text"
                    name="sympKeyword"
                    value={sympKeyword}/>

                <Options1 options={options} appendCandidate={this.appendCandidate}/>
            </div>
        )
    }
}

const mapState = state => {
    const {organs, sympAdjs} = state;
    return {organs, sympAdjs};
}

export default connect(mapState, {fillOrgans, fillSympAdjs})(Symptoms)