import { useObservable } from 'micro-observables';

import Button from 'components/kit/Button';

import css from './FilterForm.module.scss';
import { FilterFormStore } from './FilterFormStore';
import FilterGroup from './FilterGroup';
import { FormType } from './type';

interface Props {
  formStore: FilterFormStore;
}

const FilterForm = ({ formStore }: Props): JSX.Element => {
  const data = useObservable(formStore.formset);
  return (
    <div className={css.base}>
      <div className={css.filter}>
        <FilterGroup
          conjunction={data.filterGroup.conjunction}
          formStore={formStore}
          group={data.filterGroup}
          index={0}
          level={0}
          parentId={data.filterGroup.id}
        />
      </div>
      <div className={css.buttonContainer}>
        <div className={css.addButtonContainer}>
          <Button
            type="text"
            onClick={() =>
              formStore.addChild(
                data.filterGroup.id,
                FormType.Field,
                data.filterGroup.children.length,
              )
            }>
            + Add condition field
          </Button>
          <Button
            type="text"
            onClick={() =>
              formStore.addChild(
                data.filterGroup.id,
                FormType.Group,
                data.filterGroup.children.length,
              )
            }>
            + Add condition group
          </Button>
        </div>
        <Button type="text" onClick={() => formStore.removeChild(data.filterGroup.id)}>
          Clear All
        </Button>
      </div>
      <div style={{ maxWidth: '500px', wordWrap: 'break-word' }}>{formStore.query}</div>
    </div>
  );
};

export default FilterForm;
