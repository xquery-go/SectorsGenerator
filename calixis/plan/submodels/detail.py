from django.db import models
from django.forms.models import model_to_dict
from django.core.validators import int_list_validator
import json
import re

class Detail(models.Model):
    def __repr__(self):
        return json.dumps(model_to_dict(
            self,
            fields=[field.name for field in self._meta.fields]
        ))

    def __str__(self):
        return self.inspirations.all()[0].name

    def get_description(self):
        roll_list = self.rolls.split(',')
        text = '\n'.join(inspiration.description for inspiration in self.inspirations)
        for roll in roll_list:
            text = re.sub(r'\[\[[^\]]\]\]', roll, text)
        return text

    rolls = models.CharField(validators=[int_list_validator], max_length=100)
    inspirations = models.ManyToManyField('Inspiration', related_name='inspirations')
    nested_inspirations = models.ManyToManyField('Inspiration_Nested')
    parent_detail = models.ForeignKey('self', null=True, blank=True, on_delete=models.CASCADE)
