import React, { Component } from 'react';
import PropTypes from 'prop-types'
import AttackListItem from './AttackListItem'
import './List.css'

class AttackList extends Component {
    render() {
        return (
            <div className="List">
                <h3>Attacks (Ground Truth)</h3>
                <p>Drag and drop attacks into days of the timeline.</p>
                <ul>
                {Object.entries(this.props.attacks).map(entry => (
                    <AttackListItem key={entry[0]} {...entry[1]} />
                ))}
                </ul>
            </div>
        );
    }
}

AttackList.propTypes = {
    attacks: PropTypes.object.isRequired,
}

export default AttackList;