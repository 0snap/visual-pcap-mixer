import React, { Component } from 'react'
import PropTypes from 'prop-types'
import moment from 'moment';
import classNames from 'classnames';

import constants from './constants';
import styles from './List.css';


class AttackListItem extends Component {

    drag(event) {
        event.dataTransfer.setData('attackId', event.target.id);
    }

    render() {
        const cx = classNames.bind(styles);
        const classes = cx({
            disabled: this.props.added,
            enabled: !this.props.added,
        });
        const {name, attackers, victims, start, end, id} = this.props;
        return (
            <li id={id} className={classes} draggable={true} onDragStart={this.drag}>
                <span>{name}</span>
                <span>[{attackers.join(', ')}]&nbsp;&nbsp;&nbsp;{'→'}&nbsp;&nbsp;&nbsp;[{victims.join(', ')}]</span>
                <br/>
                <span>{moment(start).format(constants.TIME_FORMAT)}&nbsp;&nbsp;&nbsp;{'→'}&nbsp;&nbsp;&nbsp;{moment(end).format(constants.TIME_FORMAT)}</span>
            </li>
        );
    }
}

AttackListItem.propTypes = {
    attackers: PropTypes.arrayOf(PropTypes.string).isRequired,
    victims: PropTypes.arrayOf(PropTypes.string).isRequired,
    name: PropTypes.string.isRequired,
    start: PropTypes.string.isRequired,
    end: PropTypes.string.isRequired,
    added: PropTypes.bool.isRequired,
    id: PropTypes.number.isRequired,
}

export default AttackListItem;
