import React, { Component } from 'react';
import PropTypes from 'prop-types'
import ReplacementListItem from './ReplacementListItem'
import './List.css'

class ReplacementList extends Component {
    render() {
        return (
            <div className="List">
                <h3>Replacements</h3>
                <p>Doubleclick on an entry to delete it.</p>
                <ul>
                {this.props.replacements.map((repl, idx) => (
                    <ReplacementListItem key={idx} {...repl} onDoubleClick={() => this.props.removeReplacement(idx)}/>
                ))}
                </ul>
            </div>
        );
    }
}

ReplacementList.propTypes = {
    replacements: PropTypes.arrayOf(PropTypes.object).isRequired,
    removeReplacement: PropTypes.func.isRequired,
}

export default ReplacementList;