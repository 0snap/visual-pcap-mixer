import React, { Component } from 'react'
import PropTypes from 'prop-types'
import './List.css';

class ReplacementListItem extends Component {
    render() {
        const {ipA, ipB} = this.props;
        return (
            <li onDoubleClick={this.props.onDoubleClick}>
                <span>{ipA}</span>
                <span>{'→'}</span>
                <span>{ipB}</span>
            </li>
        );
    }
}

ReplacementListItem.propTypes = {
    ipA: PropTypes.string.isRequired,
    ipB: PropTypes.string.isRequired,
    onDoubleClick: PropTypes.func.isRequired
}

export default ReplacementListItem;