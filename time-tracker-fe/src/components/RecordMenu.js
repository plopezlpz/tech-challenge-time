import React from "react";
import { connect } from "react-redux";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import MenuItem from "@material-ui/core/MenuItem";
import Menu from "@material-ui/core/Menu";

import { RANGES } from "../constants";

import { selectRangeAndFetchRecords } from "../actions";

function RecordMenu(props) {
  const { range, selectRangeAndFetchRecords } = props;

  const [anchorEl, setAnchorEl] = React.useState(null);
  const handleClickListItem = event => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const handleMenuItemClick = rangeSelected => {
    selectRangeAndFetchRecords(rangeSelected);
    setAnchorEl(null);
  };

  return (
    <div>
      <List component="nav" aria-label="Display settings">
        <ListItem
          button
          aria-haspopup="true"
          aria-controls="range-menu"
          aria-label="select a range to display"
          onClick={handleClickListItem}
        >
          <ListItemText primary={range.label + " v"} />
        </ListItem>
      </List>
      <Menu
        id="range-menu"
        anchorEl={anchorEl}
        getContentAnchorEl={null}
        anchorOrigin={{ vertical: "bottom", horizontal: "left" }}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        {Object.keys(RANGES).map(r => (
          <MenuItem
            key={RANGES[r].id}
            selected={RANGES[r].id === range.id}
            onClick={() => handleMenuItemClick(RANGES[r])}
          >
            {RANGES[r].label}
          </MenuItem>
        ))}
      </Menu>
    </div>
  );
}

const mapStateToProps = state => {
  return { range: state.range };
};
export default connect(mapStateToProps, { selectRangeAndFetchRecords })(
  RecordMenu
);
