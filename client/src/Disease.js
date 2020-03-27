import React from "react";
import {connect} from 'react-redux'
import Symptoms from './Symptoms'
import Pathos from './Pathos'
import {redirect} from './action/page'
import config from "./helper/config";
import {fillOrgans} from "./action/organs";
import {fillSympAdjs} from "./action/sympAdjs";
import {fillOrgSyms} from "./action/orgSyms";
import {fillPathoAdjs} from './action/pathoAdjs'
import {fillPathos} from './action/pathos'
import {fillTreatVerbs} from './action/treatVerbs'
import {addDise} from './action/dises'
import fetchSelectables from "./helper/fetchSelectables";
import fetchDise from "./helper/fetchDise";

class Disease extends React.Component {
    componentDidMount = async() => {
        const {dises,addDise, fillPathos,storeOrgans, id,fillOrgans, fillOrgSyms,fillSympAdjs,fillPathoAdjs,fillTreatVerbs} = this.props;
        let dise = dises.find(x => x.id === id)
        if (!dise) {
            dise = await fetchDise(id);
            console.log(`fetch dise: `, dise)
            if (!dise){
                return alert(`after fecthing still no dise`)
            }
            addDise(dise)
        } else {
            alert(`Found local store data. todo: check if need fetch server details of disease`)
        }

        //see if organ selectables exist in store
        if (Array.isArray(storeOrgans) && storeOrgans.length > 0){
            return;
        }

        const d = await fetchSelectables();
        if (!d) 
            return alert(`fetchSelectables no d`);
        const {
            organs,
            symptomAdjs,
            organSymptoms,
            pathoAdjs,
            pathos,
            treatVerbs,
            policies,
            recipes,
            herbs
        } = d;

        console.log(`got d:`)
        console.log(d)

        //sym, patho, policy, recipe
        if (Array.isArray(organs) && organs.length > 0) {
            fillOrgans(organs);
        }

        if (Array.isArray(organSymptoms) && organSymptoms.length > 0) {
            fillOrgSyms(organSymptoms);
        }

        if (Array.isArray(symptomAdjs) && symptomAdjs.length > 0) {
            fillSympAdjs(symptomAdjs);
        }

        if (Array.isArray(pathoAdjs) && pathoAdjs.length > 0) {
            fillPathoAdjs(pathoAdjs);
        }

        if (Array.isArray(pathos) && pathos.length > 0) {
            fillPathos(pathos);
        }

        if (Array.isArray(treatVerbs) && treatVerbs.length > 0) {
            fillTreatVerbs(treatVerbs);
        }
        
    }

    render() {
        const {id, dises} = this.props;
        const dise = dises.find(x => x.id === id)
        if (!dise){
            return <div>Loading...</div>
        }

        return (
            <div>
                <div className="sectitle">Disease</div>
                <div>Name: {dise.namezh}, {dise.name}</div>
                <div>Source: {dise.source}</div>

                <Symptoms diseId={id}/>
                <Pathos diseId={id}/>
            </div>
        )
    }
}

const mapState = state => {
    const {organs, dises}=state;
    return {storeOrgans: organs, dises}
}

export default connect(mapState, {addDise,redirect, fillOrgans, fillOrgSyms,fillSympAdjs,fillPathoAdjs,fillPathos,fillTreatVerbs})(Disease)