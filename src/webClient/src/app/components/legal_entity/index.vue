<template>
  <q-page padding>
    <comp-breadcrumb :list="[{label:'ЮрЛица', docType:'legal_entity'}]"/>

    <comp-doc-list ref="docList" listTitle='ЮрЛица' listDeletedTitle='Удаленные юрЛица' pg-method="legal_entity_list"
                   :list-sort-data="listSortData" :list-filter-data="listFilterData"
                   :newDocUrl="currentUrl + 'new'"
                   search-fld-name="search_text">

      <template #listItem="{item}">
        <q-item-section avatar @click="$router.push(`${currentUrl}${item.id}`)">
          <q-avatar rounded>
            <img src="https://image.flaticon.com/icons/svg/469/469323.svg" alt="">
          </q-avatar>
        </q-item-section>
        
        <q-item-section>
          <q-item-label lines="1">{{item.title}}</q-item-label>
        </q-item-section>
	
        <q-item-section top side>
          <comp-item-dropdown-btn :item="item" itemProp="title" :is-edit="true" :is-delete="true" fkProp=""
                                  pg-method="legal_entity_update"
                                  @edit="$router.push(`${currentUrl}${item.id}`)"
                                  @reload-list="$refs.docList.reloadList()"/>
        </q-item-section>
      </template>

    </comp-doc-list>
  </q-page>
</template>

<script>
  export default {
    computed: {
      currentUrl: () => 'legal_entity/',
    },
    data() {
      return {
        listSortData: [
          {value: 'created_at', title: 'Дата'},
          {value: 'title', title: 'Название'}
        ],
        listFilterData: [
          {value: {deleted: false}, title: 'Активные'},
          {value: {deleted: true}, title: 'Удаленные'}
        ],
      }
    },
  }
</script>
