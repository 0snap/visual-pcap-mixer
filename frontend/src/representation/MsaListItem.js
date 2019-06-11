import React, { Component } from 'react'
import PropTypes from 'prop-types'

import './List.css';

class MsaListItem extends Component {
    render() {
        return (
            <li className='enabled' onClick={this.props.onClick}>
                <span>{this.props.name}</span>
            </li>
        );
    }
}

MsaListItem.propTypes = {
    name: PropTypes.string.isRequired,
    onClick: PropTypes.func.isRequired
}

export default MsaListItem;