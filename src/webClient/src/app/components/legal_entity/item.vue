<template>
  <q-page padding>

    <comp-breadcrumb :list="[{label:'ЮрЛица', to:'/legal_entity',  docType: 'legal_entity'}, {label: item ? (item.title ? item.title : 'Редактирование') : '',  docType: 'edit'}]"/>

    <div v-if="item" class="q-mt-sm">
      <!--  поля формы    -->
      
            <div class="row q-col-gutter-md q-mb-sm">
            <div class="col-4">
                <q-input outlined type='text' v-model="item.title" label="название" autogrow :readonly='false' class='col-4'/>
            </div>
            </div>
      
            <div class="row q-col-gutter-md q-mb-sm">
            <div class="col-4">
                <q-input outlined type='text' v-model="item.inn" label="ИНН" autogrow :readonly='false' class='col-4'/>
            </div>
            <div class="col-4">
                <q-input outlined type='text' v-model="item.type" label="тип организации" autogrow :readonly='false' class='col-4'/>
            </div>
            </div>
      
            <div class="row q-col-gutter-md q-mb-sm">
            <div class="col-8">
                <q-input outlined type='text' v-model="item.address_legal" label="юр адрес" autogrow :readonly='false' class='col-8'/>
            </div>
            </div>
      

      <!--  кнопки   -->
      <comp-item-btn-save @save="save" @cancel="$router.push(docUrl)"/>

    </div>
  </q-page>
</template>

<script>

    export default {
        props: ['id'],
        components: {},
        mixins: [],
        computed: {
            docUrl: () => '/legal_entity',
        },
        data() {
            return {
                item: null,
                flds: [
                        {name: 'title', label: 'название',  required: true},
                        {name: 'inn', label: 'ИНН'},
                        {name: 'kpp', label: 'КПП'},
                        {name: 'type', label: 'тип организации'},
                        {name: 'address_legal', label: 'юр адрес'},
                ],
                optionsFlds: [],
            }
        },
        methods: {
          
            resultModify(res) {
                
                return res
            },
            save() {
                this.$utils.saveItem.call(this, {
                    method: 'legal_entity_update',
                    itemForSaveMod: {},
                    resultModify: this.resultModify,
                })
            },
        },
        mounted() {
            let cb = (v) => {
                this.item = this.resultModify(v)
            }
            this.$utils.getDocItemById.call(this, {method: 'legal_entity_get_by_id', cb})
        }
    }
</script>
