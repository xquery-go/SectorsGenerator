from django.db import models
from django.forms.models import model_to_dict
import json

# Abstract Models
class BaseConfig(models.Model):
    class Meta:
        abstract = True

    def __repr__(self):
        return json.dumps(model_to_dict(
            self,
            fields=[field.name for field in self._meta.fields]
        ))

    def __str__(self):
        return self.name

    name = models.CharField(default="-", max_length=100)

class Config_Territory(BaseConfig):
    territory_inspirations = models.ManyToManyField('Weighted_Inspiration', related_name='territory_inspirations')

class Config_Element(BaseConfig):
    type_inspirations = models.ManyToManyField('Weighted_Inspiration', related_name='type_inspirations')
    spacing = models.ManyToManyField('Roll', related_name='spacing')
    satellite_count = models.ManyToManyField('Roll', related_name='satellite_count')
    satellite_extra = models.ManyToManyField('Weighted_Inspiration', related_name='satellite_extra')
    territory = models.ForeignKey('Config_Territory', null=True, blank=True, on_delete=models.SET_NULL)
    territory_count = models.ManyToManyField('Roll', related_name='territory_count')
    territory_extra = models.ManyToManyField('Weighted_Inspiration', related_name='territory_extra')

class Config_Zone(BaseConfig):
    zone = models.CharField(null=True, blank=True, max_length=25)
    distance = models.ManyToManyField('Roll', related_name='distance')
    element_count = models.ManyToManyField('Roll', related_name='element_count')
    element_extra = models.ManyToManyField('Weighted_Inspiration', related_name='element_extra')
    perterbation = models.ForeignKey('Perterbation', null=True, blank=True, on_delete=models.CASCADE)

class Config_Star_Cluster(BaseConfig):
    star_count = models.ManyToManyField('Roll', related_name='star_count')
    star_extra = models.ManyToManyField('Weighted_Inspiration', related_name='star_extra')
    star_inspirations = models.ManyToManyField('Weighted_Inspiration', related_name='star_inspirations')

class Config_Route(BaseConfig):
    stability_inspirations = models.ManyToManyField('Weighted_Inspiration', related_name='stability_inspirations')
    days_inspirations = models.ManyToManyField('Weighted_Inspiration', related_name='days_inspirations')

class Config_System(BaseConfig):
    system_feature_count = models.ManyToManyField('Roll', related_name='system_feature_count')
    system_feature_inspirations = models.ManyToManyField('Weighted_Inspiration', related_name='system_feature_inspirations')
    system_feature_extra = models.ManyToManyField('Weighted_Inspiration', related_name='system_feature_extra')
    star_cluster_count = models.ManyToManyField('Roll', related_name='star_cluster_count')

class Config_Grid(BaseConfig):
    height = models.PositiveSmallIntegerField(default=20, blank=True)
    width = models.PositiveSmallIntegerField(default=20, blank=True)
    connectionRange =models.PositiveSmallIntegerField(default=5, blank=True)
    populationRate = models.FloatField(default=0.5, blank=True)
    connectionRate = models.FloatField(default=0.5, blank=True)
    rangeRateMultiplier = models.FloatField(default=0.5, blank=True)
    smoothingFactor = models.FloatField(default=0.5, blank=True)
