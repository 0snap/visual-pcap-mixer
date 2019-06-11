import React, { Component } from 'react';
import PropTypes from 'prop-types'
import MsaTimelineDay from './MsaTimelineDay'
import './MsaTimeline.css';

class MsaTimeline extends Component {

    createDay(event) {
        event.preventDefault();
        this.props.createDay();
    }

    render() {
        const {msaTimeline, attacks, traceFiles, dayDoubleClick, entryDoubleClick, addMsaEntry} = this.props;
        const timeline = msaTimeline.map(day => day.map(entry => {
            if (entry.type === 'attack') {
                return {
                    ...attacks[entry.id],
                    ...entry,
                };
            } else if (entry.type === 'traceFile') {
                return {
                    ...traceFiles[entry.id],
                    ...entry,
                };
            }
            return {};
        }));
        
        return (
            <div className='MsaTimelinePanel'>
                <h3>Multi Step Attack Timeline</h3>
                <p>Create days in the timeline. Drag and drop days to the desired position. Hover an entry to get more information. Doubleclick to remove a day / entry.</p>
                <div onDrop={this.drop.bind(this)} onDragOver={this.allowDrop} className='MsaTimeline'>
                    {timeline.map((day, idx) => (<MsaTimelineDay key={idx} id={idx} day={day} dayDoubleClick={dayDoubleClick} entryDoubleClick={entryDoubleClick} addMsaEntry={addMsaEntry} getTraceFilesById={(id) => this.getTraceFilesById(id)}/>))}
                    <form onSubmit={this.createDay.bind(this)}>
                        <div className='FormInput'>
                            <input className='MsaTimelineAddDayBtn' type='submit' value='+'/>
                        </div>
                    </form>
                </div>
            </div>
        );
    }

    drop(event) {
        event.preventDefault();
        const draggedDayId = event.dataTransfer.getData('dayId');
        if (!draggedDayId) return;

        try {
            let lastNode = event.target;
            while (lastNode.className !== 'MsaTimelineDay') {
                lastNode = lastNode.parentNode;
            }
            const droppedOnDay = parseInt(lastNode.id, 10);
            const draggedDay = parseInt(draggedDayId, 10);
            if (!(droppedOnDay+1) || !(draggedDay+1)) {
                return
            }
            this.props.dragAndDrop(draggedDay, droppedOnDay);
        } catch (err) {
            console.error("Drag n Drop error", err);
        }
    }

    allowDrop(event) {
        event.preventDefault();
    }

    getTraceFilesById(id) {
        return this.props.traceFiles[id];
    }
}

MsaTimeline.propTypes = {
    msaTimeline: PropTypes.array.isRequired,
    attacks: PropTypes.object.isRequired,
    traceFiles: PropTypes.object.isRequired,
    dragAndDrop: PropTypes.func.isRequired,
    dayDoubleClick: PropTypes.func.isRequired,
    entryDoubleClick: PropTypes.func.isRequired,
    createDay: PropTypes.func.isRequired,
    addMsaEntry: PropTypes.func.isRequired,
}

export default MsaTimeline;