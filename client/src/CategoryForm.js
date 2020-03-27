import React from "react";
import {connect} from 'react-redux'
import saveCategory from "./helper/saveCategory"
import config from "./helper/config"
import {addOrgan} from './action/organs'
import {addSymAdj} from './action/sympAdjs'
import {addPathoAdj} from './action/pathoAdjs'
import {addTreatVerb} from './action/treatVerbs'
import {setModal} from './action/modal'

class CategoryForm extends React.Component {
    state = {
        name: ``,
        namezh: ``,
        pinyin: ``
    }

    _change = ev => {
        const {name, value} = ev.target;
        this.setState({[name]: value});
    }

    _submit = async ev => {
        ev.preventDefault();
        const {catTypeId, setModal}=this.props;
        const din = {
            catTypeId,
            ...this.state
        }
        const d = await saveCategory(din)
        console.log(`saveCategory: d` ,d)

        if (!d){
            return alert(`CateForm submit Err`)
        }

        if (!d.ok){
            return alert(`CateForm submit failed saving`)
        }

        din.id = d.id;
        din.txt=`${din.namezh} ${din.name}`
        din.searchable=`${din.txt}${din.pinyin}${din.id}`

        /*
        catTypeId: 1
        name: "mouth"
        namezh: "口"
        pinyin: "k"

        --need these
        id: 102
        frontendType: 1
        searchable: "口 mouthk102"
        txt: "口 mouth"
        
        */
       
       const {
        organ,
        sympAdj,
        pathoAdj,
        treatVerb
       } = config.categoryType;

       const {addOrgan, addSymAdj, addPathoAdj, addTreatVerb}= this.props;
        switch (catTypeId) {
            case organ:
                din.frontendType=config.frontendType.organ;
                addOrgan(din)
                break;
            case sympAdj:
                din.frontendType=config.frontendType.symptomAdj;
                addSymAdj(din)
                break;
            case pathoAdj:
                din.frontendType=config.frontendType.pathoAdj;
                addPathoAdj(din)
                break;
            case treatVerb:
                din.frontendType=config.frontendType.treatVerb;
                addTreatVerb(din)
                break;
            default:
                alert(`saved but didn't get to update store`)
                break;
        }

        setModal(false);
    }

    render() {
        const {title}=this.props;
        const {name, namezh, pinyin} = this.state;

        return <div>
            <div className="fontsize7">Add {title}</div>
            <form onSubmit={this._submit}>
                <input
                    className="ipt"
                    name="namezh"
                    placeholder="Name ZH"
                    value={namezh}
                    onChange={this._change}/>

                <input
                    className="ipt"
                    name="name"
                    placeholder="Name EN"
                    value={name}
                    onChange={this._change}/>

                <input
                    className="ipt"
                    name="pinyin"
                    placeholder="Pinyin"
                    value={pinyin}
                    onChange={this._change}/>

                <button type="submit" className="bt">Save {title}</button>
            </form>
        </div>
    }
}

const mapState2Props = state => {
    return {page: state.page, modal: state.modal}
}

export default connect(mapState2Props, {setModal, addOrgan, addSymAdj, addPathoAdj, addTreatVerb})(CategoryForm)
