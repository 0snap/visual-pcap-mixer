import { connect } from 'react-redux'
import { removeReplacement } from '../actions'
import ReplacementList from '../representation/ReplacementList';


const mapStateToProps = state => {
    return {
        replacements: state.replacements
    }
}

const mapDispatchToProps = dispatch => {
    return {
        removeReplacement: (id) => {
            dispatch(removeReplacement(id));
        }
    }
}

const ReplacementListContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(ReplacementList)

export default ReplacementListContainer