import React from "react";
import Disease from "./Disease";
import NewDisease from "./NewDisease";
import Loading from "./Loading";
import Modal from "./Modal";
import config from "./helper/config";
import urlMatch from "./helper/urlMatch";
import getIntParam from "./helper/getIntParam";
import {connect} from 'react-redux'
import {redirect} from './action/page'
import Link2 from "./Link2";
import Home from "./Home";
import CategoryForm from "./CategoryForm";
import {setModal} from './action/modal'

class App extends React.Component {
    _setModal = obj => {
        return () => {
            const {setModal} = this.props;
            setModal(obj);
        }
    }
    
    render() {
        const {redirect,modal, page} = this.props;

        if (typeof(page) !== "string") {
            return <Loading/>
        }

        const {urls} = config;

        let tail = null;

        //page urls with params
        if (urlMatch(page, urls.viewDise)) {
            const diseId = getIntParam(page, urls.viewDise);
            if (diseId > 0) {
                tail = <Disease id={diseId}/>
            }
        } else {
            //page urls without params
            switch (page) {
                case urls.newDisease:
                    tail = <NewDisease/>
                    break;
                default:
                    tail = <Home/>
            }
        }

        const {open, content}=modal;
        let kidCompo = null;
        if (open){
            const {organForm, symAdjForm,pathoAdjForm}=config.modalKid;
            if ([organForm, symAdjForm, pathoAdjForm].indexOf(content) > -1){

                let catTypeId, title;
                switch (content) {
                    case organForm:
                        catTypeId=config.categoryType.organ
                        title="organ"
                        break;
                    case symAdjForm:
                        catTypeId=config.categoryType.sympAdj
                        title="symptom adjective"
                        break;
                    case pathoAdjForm:
                        catTypeId=config.categoryType.pathoAdj
                        title="pathology adjective"
                        break;
                    default:
                        catTypeId=-1
                        title="Error"
                        break;
                }
                
                kidCompo = <CategoryForm 
                            catTypeId={catTypeId} 
                            title={title}/>
            }
        }

        return <div>
        {open && <Modal setModal={this._setModal}>
                        {kidCompo}
                    </Modal>
}

            {tail}
        </div>
    }
}

const mapState2Props = state => {
    const {page, modal}=state;
    return {page, modal}
}

export default connect(mapState2Props, {redirect, setModal})(App)