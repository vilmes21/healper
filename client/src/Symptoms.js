import React from "react";
import Options1 from "./Options1";
import Cart from "./Cart";
import Modal from "./Modal";
import CandidateTxt from "./CandidateTxt";
import CategoryForm from "./CategoryForm";
import SavedItems from "./SavedItems";
import {connect} from 'react-redux'
import {redirect} from './action/page'
import {setModal} from './action/modal'
import {addSym2Dise} from './action/dises'
import config from "./helper/config";
import saveSymps from "./helper/saveSymps";

class Symptoms extends React.Component {
    state = {
        sympKeyword: ``,
        candidate: [],
        cart: [],
        options: [],
        showOptions: false,
    }

    _focusInput = focused => {
        return () => {
            this.setState({showOptions: focused})
        }
    }

    _change = ev => {
        const {name, value} = ev.target;
        this.setState({[name]: value});

        const {organs, orgSyms, sympAdjs} = this.props;

        console.log(`onchange.  {organs, orgSyms, sympAdjs}: `)
        console.log( {organs, orgSyms, sympAdjs})
        
        const {candidate} = this.state;

        let matched;
        const collection = candidate.length === 1
            ? sympAdjs
            : [...organs, ...orgSyms, ...sympAdjs];

            //ban meaningless spaces
        const validVal = value.replace(/\s/g, "");
        if (!validVal) {
            matched = collection;
        } else {
            //organ is already set, then allow only adjective next.
            matched = collection.filter(x => x.searchable.indexOf(value) > -1);
        }

        this.setState({options: matched})
    }

    appendCandidate = x => {
        const {candidate} = this.state;
        this.setState({sympKeyword: ``})
        if (candidate.length === 2) 
            return alert(`Cannot write more. Either add to cart or modify.`)
        
        const updater = (prevState, p) => {
            const newState = {
                candidate: [
                    ...prevState.candidate,
                    x
                ],
                showOptions: false
            };

            //means organ has been set. Now only allow adjective
            if (candidate.length === 0) {
                newState.options = prevState
                    .options
                    .filter(y => y.catTypeId === config.categoryType.sympAdj);
            }

            return newState;
        }

        this.setState(updater);
    }

    rmCandidate = () => {
        this.setState({candidate: []})
    }

    cartPreBuilt = obj => {
        const {organId, symptomId, id}=obj;
        const {cart} = this.state;
        for (const item of cart) {
            if ((item.id === id) || (item.organId === organId && item.symptomId === symptomId)){
                    return alert(`Already in cart`)
                }
        }

        const updater=(prevState, p)=>{
            return {
                cart: [...prevState.cart, obj]
            }
        }
        this.setState(updater)
    }

    cartSymp = () => {
        const {cart, candidate} = this.state;
        if (candidate.length === 0) 
            return;

        if (candidate.length === 1){
            if (candidate[0].frontendType===config.frontendType.organ){
                return alert(`Must add symptom adjective. Can be just organ.`)
            } 
        } 

            let organObj = candidate.find(x => x.frontendType===config.frontendType.organ);

            if (!organObj){
                organObj = {
                    id: config.catnum.noOrganId,
                    name: ``,
                    namezh: ``
                }
            }
            
            const symObj = candidate.find(x => x.frontendType===config.frontendType.symptomAdj);

            //see if already in cart
            const duplicateInCart = cart.find(x => x.organId ===organObj.id && x.symptomId === symObj.id);
            if (duplicateInCart){
                return alert(`Duplicate in cart.`)
            }

            //see if pairing exists in db already
            const {orgSyms}=this.props;
            const prebuiltPair = orgSyms.find(x => x.organId === organObj.id && x.symptomId === symObj.id);
            
            const carting = prebuiltPair || {
                id: 0,
                organId: organObj.id,
                symptomId: symObj.id,
                txt: `${organObj.namezh}${symObj.namezh} ${organObj.name} ${symObj.name}`
            }
        
        const updater = (prevState, p) => {
            return {
                cart: [
                    ...prevState.cart,
                    carting
                ],
                candidate: []
            }
        }

        this.setState(updater);
    }

    del = arr => {
        return () => {
            console.log(`del fn,arr:`, arr)

            const updater2 = (prevState, p) => {
                return {
                    cart: prevState
                        .cart
                        .filter(x => x !== arr)
                }
            }
            this.setState(updater2);
        }
    }

    save = async() => {
        const {cart} = this.state;
        const {diseId, addSym2Dise} = this.props;

        const orgsympIds=[]
        const newSymps = []
        for (const obj of cart) {
            if (obj.id === 0){
                //see if org-sym already paired up in db
                newSymps.push({
                    organId: obj.organId,
                    symptomId: obj.symptomId
                })
            } else {
                orgsympIds.push(obj.id)
            }
        }

        const toSend = {
            diseId,
            orgsympIds,
            newSymps
        }

        const d = await saveSymps(toSend);
        if (d && d.ok){
            addSym2Dise({
                diseId,
                syms: d.diseSyms
            });
            this.setState({cart: []})
        } else {
            return alert(`Err saving symptoms`)
        }
    }

    _setModal = (open, content) => {
        const {setModal}=this.props;
        return () =>{
            setModal({open, content});
        }
    }

    render() {
        const {sympKeyword, cart, showOptions, candidate, options} = this.state;

        const {organForm, symAdjForm}=config.modalKid
        const {diseId, dises} = this.props;
        const diseNow = dises.find(x => x.id === this.props.diseId);
        const savedDiseSyms = diseNow && Array.isArray(diseNow.syms)? diseNow.syms : [];

        return (
            <div>
                <div className="sectitle">Symptoms</div>

                <SavedItems items={savedDiseSyms} ownerId={diseId}/>

                <Cart 
                itemTypeStr={"symptoms"}
                save={this.save} objArr={cart} del={this.del}/>

                <div>
                    <CandidateTxt objArr={candidate}/> {candidate.length > 0 && <div>
                        <button className="bt" onClick={this.cartSymp}>Add to cart</button>

                        <button className="bt" onClick={this.rmCandidate}>Remove</button>
                    </div>
}
                </div>

                <input
                autoComplete={"off"}
                    className="ipt"
                    onFocus={this._focusInput(true)}
                    onChange={this._change}
                    type="text"
                    name="sympKeyword"
                    value={sympKeyword}/> 
                
                <button onClick={this._setModal(true, organForm)}>Add organ</button>
                <button onClick={this._setModal(true, symAdjForm)}
                >Add symptom adjective</button>
                    
                    {showOptions && <Options1 options={options} appendCandidate={this.appendCandidate}
                    cartPreBuilt={this.cartPreBuilt}/>
}
            </div>
        )
    }
}

const mapState = state => {
    const {organs, orgSyms, sympAdjs, dises} = state;
    return {organs, orgSyms, sympAdjs, dises};
}

export default connect(mapState, {setModal, addSym2Dise})(Symptoms)