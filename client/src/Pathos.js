import React from "react";
import Options1 from "./Options1";
import Cart from "./Cart";
import CandidateTxt from "./CandidateTxt";
import SavedItems from "./SavedItems";
import {connect} from 'react-redux'
import {redirect} from './action/page'
import {setModal} from './action/modal'
import {addSym2Dise} from './action/dises'
import config from "./helper/config";
import saveSymps from "./helper/saveSymps";

class Pathos extends React.Component {
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
    
        const {organs, pathos, pathoAdjs} = this.props;
        const {candidate} = this.state;
    
        let matched;
    
        //organ & tcm_orang are already set, then allow only adjective next.
        const collection = candidate.length === 2
            ? pathoAdjs
            : [...pathos, ...organs, ...pathoAdjs];
    
            //ban meaningless spaces
        const validVal = value.replace(/\s/g, "");
        if (!validVal) {
            matched = collection;
        } else {
            
            matched = collection.filter(x => x.searchable.indexOf(value) > -1);
        }
    
        this.setState({options: matched})
    }

    appendCandidate = x => {
        const {candidate} = this.state;
        this.setState({sympKeyword: ``})
        if (candidate.length === 3) 
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
            if (candidate.length === 1) {
                newState.options = prevState
                    .options
                    .filter(y => y.catTypeId === config.categoryType.pathoAdj);
            }

            return newState;
        }

        this.setState(updater);
    }

    rmCandidate = () => {
        this.setState({candidate: []})
    }

    cartPreBuilt = obj => {
        const {organId,tcmOrganId, descriptionId, id}=obj;
        const {cart} = this.state;
        for (const item of cart) {
            if ((item.id === id) || (item.organId === organId 
                && item.tcmOrganId === tcmOrganId  
                && item.descriptionId === descriptionId)){
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

    cartFresh = () => {
        const {cart, candidate} = this.state;
        if (candidate.length === 0) 
            return;

        const feOrganType = config.frontendType.organ;
        const feAdj = config.frontendType.pathoAdj;
       
        let organObj = {
            id: config.catnum.noOrganId,
            name: ``,
            namezh: ``
        }

        let tcmOrganObj={
            id: config.catnum.noOrganId,
            name: ``,
            namezh: ``
        }
        
        if (candidate.length === 1){
            if (candidate[0].frontendType!==feAdj){
                return alert(`Must add pathology adjective. Can't be just organ.`)
            }
        } else if (candidate.length === 2){
            if (candidate[0].frontendType===feOrganType && candidate[1].frontendType===feOrganType){
                return alert(`Must add pathology adjective. Can't be just organ.`)
            } 

            if (candidate[0].frontendType===feAdj){
                return alert(`Must place pathology adjective before organ.`)
            }

            tcmOrganObj = candidate[0];
        } else if (candidate.length === 3){
            //must not have more than 1 adj
            let adjCount= 0;
            for (const obj of candidate) {
                if (obj.frontendType===feAdj){
                    if (adjCount === 1){
                        return alert(`Can use only 1 adjective`)
                    }
                    adjCount++;
                }
            }

            organObj = candidate[0];
            tcmOrganObj = candidate[1];
        }

        const adjObj = candidate.find(x => x.frontendType===feAdj);

        //see if already in cart
        const duplicateInCart = cart.find(x => 
            x.organId ===organObj.id 
            && x.tcmOrganId === tcmOrganObj.id
            && x.descriptionId === adjObj.id
            );
        
        if (duplicateInCart){
            return alert(`Duplicate in cart.`)
        }
        
        //see if pairing exists in db already
        const {pathos}=this.props;
        
        /*
        from server searchable:
            id: 1
            frontendType: 4
            organId: 111
            tcmOrganId: 116
            descriptionId: 501
            txt: "耳阳虚 ears-yang-deficient"
            searchable: "耳阳虚 ears-""
        
        */
        const prebuiltPair = pathos.find(x => 
            x.organId ===organObj.id 
            && x.tcmOrganId === tcmOrganObj.id
            && x.descriptionId === adjObj.id
            );
        
        //`carting` must contain key `descriptionId` because other server-supplied objs use that key `descriptionId`
        const carting = prebuiltPair || {
            id: 0,
            organId: organObj.id,
            tcmOrganId: tcmOrganObj.id,
            descriptionId: adjObj.id,
            txt: `${organObj.namezh}${tcmOrganObj.namezh}${adjObj.namezh} ${organObj.name} ${tcmOrganObj.name} ${adjObj.name}`
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

        this.setState(updater, () => {
            console.log(`after patho car:`, this.state)
        });
    }

    del = arr => {
        return () => {
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
        console.log(`patho save: d`, d)
        
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

        const {pathoAdjForm}=config.modalKid
        const {diseId, dises} = this.props;
        const diseNow = dises.find(x => x.id === this.props.diseId);
        const savedDisePathos = diseNow && Array.isArray(diseNow.pathos)? diseNow.pathos : [];

        return (
            <div>
                <div className="sectitle">Pathology</div>

                <SavedItems items={savedDisePathos} ownerId={diseId}/>

                <Cart
                itemTypeStr={"pathology"}
                save={this.save} objArr={cart} del={this.del}/>

                <div>
                    <CandidateTxt objArr={candidate}/> {candidate.length > 0 && <div>
                        <button className="bt" onClick={this.cartFresh}>Add to cart</button>

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
                
                <button onClick={this._setModal(true, pathoAdjForm)}
                >Add pathology adjective</button>
                    
                    {showOptions && <Options1 
                                        options={options} 
                                        appendCandidate={this.appendCandidate}
                                        cartPreBuilt={this.cartPreBuilt}/>
}
            </div>
        )
    }
}

const mapState = state => {
    const {organs, pathos, pathoAdjs,dises} = state;
    return {organs, pathos, pathoAdjs,dises};
}

export default connect(mapState, {setModal, addSym2Dise})(Pathos)