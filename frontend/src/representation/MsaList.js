import React, { Component } from 'react';
import PropTypes from 'prop-types'
import MsaListItem from './MsaListItem'
import './List.css'

class MsaList extends Component {
    render() {
        const allMsa = this.props.multistepattacks;
        return (
            <div className="List">
                <h3>Multi Step Attacks</h3>
                <p>Load some existing MSA to the timeline by clicking on it.</p>
                <ul>
                {Object.keys(allMsa).map((name, idx) => (
                    <MsaListItem key={idx} name={name} onClick={() => this.props.loadMsaToTimeline(allMsa[name])}/>
                ))}
                </ul>
            </div>
        );
    }
}

MsaList.propTypes = {
    multistepattacks: PropTypes.object.isRequired,
    loadMsaToTimeline: PropTypes.func.isRequired,
}

export default MsaList;